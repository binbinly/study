package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"

	"chat/app/logic/idl"
	"chat/app/logic/model"
	"chat/app/message"
	"chat/pkg/app"
	"chat/pkg/log"
)

const (
	//MsgCreate 创建群聊消息
	MsgCreate = "群聊已创建，可以开始聊天啦"
	//MsgEditName 修改群名
	MsgEditName = "修改群聊名为 %s"
	//MsgEditRemark 修改公告
	MsgEditRemark = "[新公告] %s"
	//MsgKickoff 踢成员
	MsgKickoff = "将 %s 移出了群聊"
	//MsgInvite 邀请好友
	MsgInvite = "邀请 %s 加入了群聊"
	//MsgJoin 加入群
	MsgJoin = "加入了群聊"
	//MsgQuit 退出群
	MsgQuit = "退出了该群聊"
	//MsgDisband 解散群
	MsgDisband = "解散了群聊"
)

// 发送消息结构体
type sendParams struct {
	userID   uint32                  // 操作人ID
	group    *model.GroupModel       // 群模型
	gUsers   []*model.GroupUserModel // 群成员模型数组
	content  string                  // 推送消息内容
	targetID uint32                  // 目标人
	tContent string                  // 目标人消息
}

//IGroup 群组服务接口
type IGroup interface {
	// 创建群组
	GroupCreate(ctx context.Context, userID uint32, ids []uint32) error
	// 修改群组名
	GroupEditName(ctx context.Context, userID, groupID uint32, name string) error
	// 修噶群公告
	GroupEditRemark(ctx context.Context, userID, groupID uint32, remark string) error
	// 修改我的群昵称
	GroupEditUserNickname(ctx context.Context, userID, groupID uint32, nickname string) error
	// 我的群列表
	GroupMyList(ctx context.Context, userID uint32) (list []*model.GroupList, err error)
	// 群详情
	GroupInfo(ctx context.Context, userID, groupID uint32) (info *model.GroupInfo, err error)
	// 群成员
	GroupUserAll(ctx context.Context, userID, groupID uint32) (list []*model.UserBase, err error)
	// 退出群
	GroupUserQuit(ctx context.Context, userID, groupID uint32) (err error)
	// 踢出群
	GroupKickOffUser(ctx context.Context, myID, groupID, toID uint32) (err error)
	// 邀请入群
	GroupInviteUser(ctx context.Context, myID, groupID, toID uint32) (err error)
	// 加入群
	GroupJoin(ctx context.Context, userID, groupID uint32) (err error)
}

// GroupCreate 创建群组
func (s *Service) GroupCreate(ctx context.Context, userID uint32, ids []uint32) error {
	u, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return errors.Wrapf(err, "[service.group] user")
	}
	uname := idl.TransferUserName(u)
	// 选择好友信息
	friends, err := s.repo.GetFriendsByIds(ctx, userID, ids)
	if err != nil {
		return errors.Wrapf(err, "[service.group] friends err")
	}
	if len(friends) == 0 {
		return ErrFriendNotRecord
	}
	// 好友id列表
	userIDs := make([]uint32, 0)
	for _, f := range friends {
		userIDs = append(userIDs, f.FriendID)
	}
	// 批量获取用户信息
	users, err := s.repo.GetUsersByIds(ctx, userIDs)
	if err != nil {
		return errors.Wrapf(err, "[service.friend] get user ids:%v", userIDs)
	}
	group := &model.GroupModel{
		UID:  model.UID{UserID: userID},
		Name: s.groupName(uname, users),
	}
	// 开启事务
	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 创建群组
	groupID, err := s.repo.GroupCreate(ctx, tx, group)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.group] create")
	}
	// 创建群组成员
	err = s.repo.GroupUserBatchCreate(ctx, tx, userID, groupID, users)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.group] users create")
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.group] tx commit err")
	}
	from := &message.From{
		ID:     u.ID,
		Name:   uname,
		Avatar: u.Avatar,
	}
	to := &message.From{
		ID:     group.ID,
		Name:   group.Name,
		Avatar: group.Avatar,
	}
	now := time.Now().Unix()
	// 给创建者发送消息
	msg, err := app.NewMessagePack(message.EventChat, &message.Chat{
		From:     from,
		To:       to,
		ChatType: model.MessageChatTypeGroup,
		Type:     model.MessageTypeSystem,
		Content:  MsgCreate,
		T:        now,
	})
	if err != nil {
		return err
	}
	uids := []uint32{userID}
	// 给好友发送消息
	for _, f := range friends {
		uids = append(uids, f.FriendID)
	}
	// 发送消息
	err = s.PushUserIds(ctx, uids, message.EventChat, msg)
	if err != nil {
		return errors.Wrapf(err, "[service.group] send")
	}
	return nil
}

// GroupEditName 更新群名
func (s *Service) GroupEditName(ctx context.Context, userID, groupID uint32, name string) error {
	// 群信息
	group, gUsers, err := s.groupInfo(ctx, groupID)
	if err != nil {
		return err
	}
	_, my := s.groupUserIds(userID, gUsers)
	if my == nil {
		return ErrGroupUserNotJoin
	}
	if group.Name == name { // 群名没变，无需修改
		return ErrGroupDataUnmodified
	}
	group.Name = name
	// 修改群组信息
	err = s.repo.GroupSave(ctx, group)
	if err != nil {
		return errors.Wrapf(err, "[service.group] save by id: %d", groupID)
	}
	return s.sendMessage(ctx, &sendParams{
		userID:  userID,
		group:   group,
		gUsers:  gUsers,
		content: fmt.Sprintf(MsgEditName, name),
	})
}

// GroupEditRemark 更新群公告
func (s *Service) GroupEditRemark(ctx context.Context, userID, groupID uint32, remark string) error {
	// 群信息
	group, gUsers, err := s.groupInfo(ctx, groupID)
	if err != nil {
		return err
	}
	_, my := s.groupUserIds(userID, gUsers)
	if my == nil {
		return ErrGroupUserNotJoin
	}
	if group.Remark == remark { // 公告没变，无需修改
		return ErrGroupDataUnmodified
	}
	group.Remark = remark
	// 修改群组信息
	err = s.repo.GroupSave(ctx, group)
	if err != nil {
		return errors.Wrapf(err, "[service.group] save group by id: %d", groupID)
	}
	return s.sendMessage(ctx, &sendParams{
		userID:  userID,
		group:   group,
		gUsers:  gUsers,
		content: fmt.Sprintf(MsgEditRemark, remark),
	})
}

// GroupEditUserNickname 更新我在群组中的昵称
func (s *Service) GroupEditUserNickname(ctx context.Context, userID, groupID uint32, nickname string) error {
	// 群信息
	_, gUsers, err := s.groupInfo(ctx, groupID)
	if err != nil {
		return err
	}
	_, my := s.groupUserIds(userID, gUsers)
	if my == nil {
		return ErrGroupUserNotJoin
	}
	return s.repo.GroupUserUpdateNickname(ctx, userID, groupID, nickname)
}

// GroupMyList 我的群组
func (s *Service) GroupMyList(ctx context.Context, userID uint32) (list []*model.GroupList, err error) {
	list, err = s.repo.GetGroupsByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.group] MyGroupList uid:%d", userID)
	}
	return
}

// GroupInfo 群组信息
func (s *Service) GroupInfo(ctx context.Context, userID, groupID uint32) (info *model.GroupInfo, err error) {
	// 群信息
	group, gUsers, err := s.groupInfo(ctx, groupID)
	if err != nil {
		return nil, err
	}
	userIDs, my := s.groupUserIds(userID, gUsers)
	if my == nil {
		return nil, ErrGroupUserNotJoin
	}
	// 批量获取用户信息
	users, err := s.repo.GetUsersByIds(ctx, userIDs)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.group] users ids:%v", userIDs)
	}
	input := &idl.TransferGroupInput{
		Group:     group,
		GroupUser: gUsers,
		Users:     users,
		Nickname:  my.Nickname,
	}
	return idl.TransferGroupInfo(input), nil
}

// GroupUserAll 所有群成员
func (s *Service) GroupUserAll(ctx context.Context, userID, groupID uint32) (list []*model.UserBase, err error) {
	// 群信息
	_, gUsers, err := s.groupInfo(ctx, groupID)
	if err != nil {
		return nil, err
	}
	userIDs, my := s.groupUserIds(userID, gUsers)
	if my == nil {
		return nil, ErrGroupUserNotJoin
	}
	// 批量获取用户信息
	users, err := s.repo.GetUsersByIds(ctx, userIDs)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.group] users ids:%v", userIDs)
	}
	input := &idl.TransferGroupInput{
		GroupUser: gUsers,
		Users:     users,
	}
	return idl.TransferGroupUser(input), nil
}

// GroupUserQuit 删除并退出群聊
func (s *Service) GroupUserQuit(ctx context.Context, userID, groupID uint32) (err error) {
	// 群信息
	group, gUsers, err := s.groupInfo(ctx, groupID)
	if err != nil {
		return err
	}
	if group.UserID == userID { // 管理员解散群
		return s.deleteGroup(ctx, userID, group, gUsers)
	}
	// 删除群成员
	return s.deleteGroupUser(ctx, userID, group, gUsers)
}

// GroupKickOffUser 踢出群成员
func (s *Service) GroupKickOffUser(ctx context.Context, myID, groupID, toID uint32) (err error) {
	// 群信息
	group, gUsers, err := s.groupInfo(ctx, groupID)
	if err != nil {
		return err
	}
	_, my := s.groupUserIds(myID, gUsers)
	if my == nil {
		return ErrGroupUserNotJoin
	}
	// 目标用户信息
	toUser, err := s.repo.GetUserByID(ctx, toID)
	if err != nil {
		return errors.Wrapf(err, "[service.group] user id:%d", toID)
	}
	if toUser.ID == 0 {
		return ErrUserNotFound
	}
	_, to := s.groupUserIds(toID, gUsers)
	if to == nil {
		return ErrGroupUserTargetNotJoin
	}
	// 被踢人昵称
	kName := to.Nickname
	err = s.repo.GroupUserDelete(ctx, to)
	if err != nil {
		return errors.Wrapf(err, "[service.group] kickoff err uid:%d,gid:%d", toID, groupID)
	}
	if kName == "" {
		if toUser.Nickname != "" {
			kName = toUser.Nickname
		} else {
			kName = toUser.Username
		}
	}
	return s.sendMessage(ctx, &sendParams{
		userID:   myID,
		group:    group,
		gUsers:   gUsers,
		content:  fmt.Sprintf(MsgKickoff, kName),
		targetID: toID,
		tContent: fmt.Sprintf(MsgKickoff, "你"),
	})
}

// GroupInviteUser 邀请好友加入
func (s *Service) GroupInviteUser(ctx context.Context, myID, groupID, toID uint32) (err error) {
	// 群信息
	group, gUsers, err := s.groupInfo(ctx, groupID)
	if err != nil {
		return err
	}
	_, my := s.groupUserIds(myID, gUsers)
	if my == nil {
		return ErrGroupUserNotJoin
	}
	// 目标用户信息
	toUser, err := s.repo.GetUserByID(ctx, toID)
	if err != nil {
		return errors.Wrapf(err, "[service.group] user id:%d", toID)
	}
	if toUser.ID == 0 {
		return ErrUserNotFound
	}
	_, to := s.groupUserIds(toID, gUsers)
	if to != nil {
		return ErrGroupUserExisted
	}
	// 加入群聊
	gUser := &model.GroupUserModel{
		UID:     model.UID{UserID: toID},
		GroupID: group.ID,
	}
	err = s.repo.GroupUserCreate(ctx, gUser)
	if err != nil {
		return errors.Wrapf(err, "[service.group] CreateUser uid:%d, gid:%d", toID, groupID)
	}
	name := toUser.Username
	if toUser.Nickname != "" {
		name = toUser.Nickname
	}
	gUsers = append(gUsers, gUser)
	return s.sendMessage(ctx, &sendParams{
		userID:   myID,
		group:    group,
		gUsers:   gUsers,
		content:  fmt.Sprintf(MsgInvite, name),
		targetID: toID,
		tContent: fmt.Sprintf(MsgInvite, "你"),
	})
}

// GroupJoin 加入群
func (s *Service) GroupJoin(ctx context.Context, userID, groupID uint32) (err error) {
	// 群信息
	group, gUsers, err := s.groupInfo(ctx, groupID)
	if err != nil {
		return err
	}
	_, my := s.groupUserIds(userID, gUsers)
	if my != nil {
		return ErrGroupUserExisted
	}
	// 加入群聊
	gUser := &model.GroupUserModel{
		UID:     model.UID{UserID: userID},
		GroupID: group.ID,
	}
	err = s.repo.GroupUserCreate(ctx, gUser)
	if err != nil {
		return errors.Wrapf(err, "[service.group] CreateUser uid:%d, gid:%d", userID, groupID)
	}
	gUsers = append(gUsers, gUser)
	return s.sendMessage(ctx, &sendParams{
		userID:  userID,
		group:   group,
		gUsers:  gUsers,
		content: MsgJoin,
	})
}

// groupName 获取群组默认名称
func (s *Service) groupName(username string, friends []*model.UserModel) string {
	var name strings.Builder
	name.WriteString(username)
	for i, f := range friends {
		if i == 4 { //最多拼接4位好友昵称
			break
		}
		m := f.Username
		if f.Nickname != "" {
			m = f.Nickname
		}
		name.WriteString(",")
		name.WriteString(m)
	}
	return name.String()
}

// sendMessage 发送群消息
func (s *Service) sendMessage(ctx context.Context, params *sendParams) (err error) {
	mContent := params.content
	// 我的用户详情
	u, err := s.repo.GetUserByID(ctx, params.userID)
	if err != nil {
		return errors.Wrapf(err, "[service.group] user info id:%v", params.userID)
	}
	params.content = fmt.Sprintf("%s %s", s.myGroupName(u, params.gUsers), params.content)

	f := &message.From{
		ID:     u.ID,
		Name:   s.myGroupName(u, params.gUsers),
		Avatar: u.Avatar,
	}
	t := &message.From{
		ID:     params.group.ID,
		Name:   params.group.Name,
		Avatar: params.group.Avatar,
	}
	now := time.Now().Unix()
	m := make([]*PushReq, 0)
	// 给群组成员发送消息
	for _, gUser := range params.gUsers {
		c := params.content
		if gUser.UserID == params.userID { // 发送给自己的消息
			c = "你 " + mContent
		}
		if gUser.UserID == params.targetID {
			c = fmt.Sprintf("%s %s", s.myGroupName(u, params.gUsers), params.tContent)
		}
		msg, err := app.NewMessagePack(message.EventChat, &message.Chat{
			From:     f,
			To:       t,
			ChatType: model.MessageChatTypeGroup,
			Type:     model.MessageTypeSystem,
			Content:  c,
			T:        now,
		})
		if err != nil {
			log.Warn(err)
			continue
		}
		m = append(m, &PushReq{
			UserID: gUser.UserID,
			Event:  message.EventChat,
			Data:   msg,
		})
	}
	// 推送消息
	err = s.PushBatch(ctx, m)
	if err != nil {
		return errors.Wrapf(err, "[service.group] send")
	}
	return nil
}

// myGroupName 我再群组中显示的昵称
func (s *Service) myGroupName(my *model.UserModel, users []*model.GroupUserModel) string {
	name := my.Username
	if my.Nickname != "" {
		name = my.Nickname
	}
	for _, u := range users { // 获取我再群组中的昵称
		if u.UserID == my.ID {
			if u.Nickname != "" {
				name = u.Nickname
			}
			break
		}
	}
	return name
}

// 获取群信息
func (s *Service) groupInfo(ctx context.Context, groupID uint32) (group *model.GroupModel, gUsers []*model.GroupUserModel, err error) {
	group, err = s.repo.GetGroupByID(ctx, groupID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[service.group] info id:%d", groupID)
	}
	if group.ID == 0 {
		return nil, nil, ErrGroupNotFound
	}
	// 群组成员
	gUsers, err = s.repo.GroupUserAll(ctx, groupID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[service.group] user list id:%d", groupID)
	}
	return
}

// deleteGroupUser 删除群成员
func (s *Service) deleteGroupUser(ctx context.Context, userID uint32, group *model.GroupModel, gUsers []*model.GroupUserModel) (err error) {
	_, my := s.groupUserIds(userID, gUsers)
	if my == nil {
		return ErrGroupUserNotJoin
	}
	err = s.repo.GroupUserDelete(ctx, my)
	if err != nil {
		return errors.Wrapf(err, "[service.group] quit err uid:%d,gid:%d", userID, my.GroupID)
	}
	return s.sendMessage(ctx, &sendParams{
		userID:  userID,
		group:   group,
		gUsers:  gUsers,
		content: MsgQuit,
	})
}

// deleteGroup 删除群组
func (s *Service) deleteGroup(ctx context.Context, userID uint32, group *model.GroupModel, gUsers []*model.GroupUserModel) (err error) {
	// 开启事务
	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 删除群
	err = s.repo.GroupDelete(ctx, tx, group)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.group] delete err")
	}
	// 删除群成员
	err = s.repo.GroupUserDeleteByGroupID(ctx, tx, group.ID)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.group] delete users err")
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "tx commit err")
	}
	return s.sendMessage(ctx, &sendParams{
		userID:  userID,
		group:   group,
		gUsers:  gUsers,
		content: MsgDisband,
	})
}

// groupUserIds 获取群信息
func (s *Service) groupUserIds(userID uint32, users []*model.GroupUserModel) (ids []uint32, my *model.GroupUserModel) {
	ids = make([]uint32, 0)
	for _, u := range users {
		ids = append(ids, u.UserID)
		if u.UserID == userID {
			my = u
		}
	}
	return
}
