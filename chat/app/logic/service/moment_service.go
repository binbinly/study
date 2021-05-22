package service

import (
	"context"

	"github.com/pkg/errors"

	"chat/app/logic/idl"
	"chat/app/logic/message"
	"chat/app/logic/model"
	"chat/app/constvar"
	"chat/pkg/app"
	"chat/pkg/utils"
)

type IMoment interface {
	// 发布朋友圈
	MomentPush(ctx context.Context, userId uint32, content, image, video, location string, t, sType int8, remind, see []uint32) (err error)
	// 我的朋友圈
	MomentTimeline(ctx context.Context, userId uint32, offset int) (*model.Moment, error)
	// 好友朋友圈
	MomentList(ctx context.Context, myId, userId uint32, offset int) (*model.Moment, error)
	// 点赞
	MomentLike(ctx context.Context, userId, momentId uint32) error
	// 评论
	MomentComment(ctx context.Context, userId, replyId, momentId uint32, content string) error
}

// MomentPush 发布朋友圈
func (s *Service) MomentPush(ctx context.Context, userId uint32, content, image, video, location string, t, sType int8, remind, see []uint32) (err error) {
	// 我的好友列表
	friends, err := s.repo.GetFriendAll(ctx, userId)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] friends err")
	}
	// 好友id列表
	friendIds := make([]uint32, 0)
	for _, f := range friends {
		friendIds = append(friendIds, f.FriendId)
	}
	// 过滤非好友元素
	newRemind := utils.FilterSmallUInt32Slice(friendIds, func(v uint32) bool {
		return utils.InuInt32Slice(v, remind)
	})
	m := &model.MomentModel{
		Uid:      model.Uid{UserId: userId},
		Content:  content,
		Image:    image,
		Video:    video,
		Location: location,
		Remind:   utils.SliceUInt32ToString2(newRemind),
		Type:     t,
		SeeType:  sType,
		See:      utils.SliceUInt32ToString2(see),
	}
	// 开启事务
	db := model.GetDB()
	tx := db.Begin()
	id, err := s.repo.MomentCreate(ctx, tx, m)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.moment] create err")
	}
	// 时间线
	lines := make([]*model.MomentTimelineModel, 0)
	// 自己
	lines = append(lines, &model.MomentTimelineModel{
		Uid:      model.Uid{UserId: userId},
		MomentId: id,
		IsOwn:    1,
	})
	for _, f := range friends {
		line := &model.MomentTimelineModel{
			Uid:      model.Uid{UserId: f.FriendId},
			MomentId: id,
			IsOwn:    0,
		}
		if sType == model.MomentSeeTypeAll {
			lines = append(lines, line)
		} else if sType == model.MomentSeeTypeOnly {
			if utils.InuInt32Slice(f.FriendId, see) {
				lines = append(lines, line)
			}
		} else if sType == model.MomentSeeTypeExcept {
			if !utils.InuInt32Slice(f.FriendId, see) {
				lines = append(lines, line)
			}
		}
	}
	_, err = s.repo.TimelineBatchCreate(ctx, tx, lines)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.moment] timeline batch create err")
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.moment] tx commit err")
	}
	return s.pushMessage(ctx, userId, lines, newRemind)
}

// MomentTimeline 我的朋友圈动态
func (s *Service) MomentTimeline(ctx context.Context, userId uint32, offset int) (*model.Moment, error) {
	u, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] user err id:%d", userId)
	}
	if u.ID == 0 {
		return nil, ErrUserNotFound
	}
	// 朋友圈动态
	mList, err := s.repo.GetMyMoments(ctx, userId, offset, constvar.DefaultLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] list err uid:%d", userId)
	}
	if len(mList) == 0 {
		return &model.Moment{
			User: idl.TransferUserBase(u),
			List: make([]*model.MomentList, 0),
		}, nil
	}
	mIds := make([]uint32, 0)
	// 先用map存放，为去重用户id
	mapUserIds := make(map[uint32]bool, 0)
	for _, momentModel := range mList {
		mIds = append(mIds, momentModel.ID)
		mapUserIds[momentModel.UserId] = true
	}
	// 点赞信息
	likeList, err := s.repo.GetLikesByMomentIds(ctx, mIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] likes err mids:%v", mIds)
	}
	for _, like := range likeList {
		for _, uid := range *like {
			mapUserIds[uid] = true
		}
	}
	// 评论信息
	commentList, err := s.repo.GetCommentsByMomentIds(ctx, mIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] comments err mids:%v", mIds)
	}
	for _, list := range commentList {
		for _, commentModel := range *list {
			mapUserIds[commentModel.UserId] = true
			if commentModel.ReplyId > 0 {
				mapUserIds[commentModel.ReplyId] = true
			}
		}
	}
	// 汇总所有动态，点赞，评论，恢复的用户id
	userIds := make([]uint32, 0)
	for uid := range mapUserIds {
		userIds = append(userIds, uid)
	}
	// 批量获取用户信息
	users, err := s.repo.GetUsersByIds(ctx, userIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] users err")
	}
	input := &idl.TransferMomentInput{
		Moments:     mList,
		Users:       users,
		LikeList:    likeList,
		CommentList: commentList,
	}
	return &model.Moment{
		User: idl.TransferUserBase(u),
		List: idl.TransferMomentList(input),
	}, nil
}

// MomentList 指定好友的动态
func (s *Service) MomentList(ctx context.Context, myId, userId uint32, offset int) (*model.Moment, error) {
	u, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] user err id:%d", userId)
	}
	if u.ID == 0 {
		return nil, ErrUserNotFound
	}
	// 朋友圈动态
	mList, err := s.repo.GetMomentsByUserId(ctx, myId, userId, offset, constvar.DefaultLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] list err uid:%d", userId)
	}
	if len(mList) == 0 {
		return &model.Moment{
			User: idl.TransferUserBase(u),
			List: make([]*model.MomentList, 0),
		}, nil
	}
	// 我的好友列表
	friends, err := s.repo.GetFriendAll(ctx, myId)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] friends err")
	}
	// 好友id列表
	friendIds := make([]uint32, 0)
	for _, f := range friends {
		friendIds = append(friendIds, f.FriendId)
	}
	mIds := make([]uint32, 0)
	// 先用map存放，为去重用户id
	mapUserIds := make(map[uint32]bool, 0)
	for _, momentModel := range mList {
		mIds = append(mIds, momentModel.ID)
		mapUserIds[momentModel.UserId] = true
	}
	// 点赞信息
	likeList, err := s.repo.GetLikesByMomentIds(ctx, mIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] likes err mids:%v", mIds)
	}
	for _, like := range likeList {
		for _, uid := range *like {
			if utils.InuInt32Slice(uid, friendIds) { // 过滤非我的好友的点赞
				mapUserIds[uid] = true
			}
		}
	}
	// 评论信息
	commentList, err := s.repo.GetCommentsByMomentIds(ctx, mIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] comments err mids:%v", mIds)
	}
	for _, list := range commentList {
		for _, commentModel := range *list {
			if utils.InuInt32Slice(commentModel.UserId, friendIds) { // 过滤非我的好友的评论
				mapUserIds[commentModel.UserId] = true
				if commentModel.ReplyId > 0 {
					if utils.InuInt32Slice(commentModel.ReplyId, friendIds) {
						mapUserIds[commentModel.ReplyId] = true
					}
				}
			}
		}
	}
	// 汇总所有动态，点赞，评论，恢复的用户id
	userIds := make([]uint32, 0)
	for uid := range mapUserIds {
		userIds = append(userIds, uid)
	}
	// 批量获取用户信息
	users, err := s.repo.GetUsersByIds(ctx, userIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] users err")
	}
	input := &idl.TransferMomentInput{
		Moments:     mList,
		Users:       users,
		LikeList:    likeList,
		CommentList: commentList,
	}
	return &model.Moment{
		User: idl.TransferUserBase(u),
		List: idl.TransferMomentList(input),
	}, nil
}

// MomentLike 点赞
func (s *Service) MomentLike(ctx context.Context, userId, momentId uint32) error {
	u, authorId, err := s.momentCheck(ctx, userId, momentId)
	if err != nil {
		return err
	}
	// 已经点赞的用户列表
	likeIds, err := s.repo.GetLikeUserIdsByMomentId(ctx, momentId)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] userIds err mid:%d", momentId)
	}
	// 是否已点赞
	isLike, err := s.repo.LikeExist(ctx, userId, momentId)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] exist like err uid:%d mid:%d", userId, momentId)
	}
	if isLike { // 已点赞，即取消
		err = s.repo.LikeDelete(ctx, userId, momentId)
		if err != nil {
			return errors.Wrapf(err, "[service.moment] delete err uid:%d mid:%d", userId, momentId)
		}
		return nil
	}
	// 创建点赞记录
	mLike := &model.MomentLikeModel{
		UserId:   userId,
		MomentId: momentId,
	}
	_, err = s.repo.LikeCreate(ctx, mLike)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] create like err uid:%d mid:%d", userId, momentId)
	}
	// 通知作者
	msg, err := app.NewMessagePack(message.EventMoment, &message.Moment{
		UserId: userId,
		Avatar: u.Avatar,
		Type:   "like",
	})
	if err != nil {
		return err
	}
	userIds := []uint32{authorId}
	// 发送其他点赞好友
	for _, id := range *likeIds {
		userIds = append(userIds, id)
	}
	return s.PushUserIds(ctx, userIds, message.EventMoment, msg)
}

// MomentComment 评论
func (s *Service) MomentComment(ctx context.Context, userId, replyId, momentId uint32, content string) error {
	u, authorId, err := s.momentCheck(ctx, userId, momentId)
	if err != nil {
		return err
	}
	// 已评论用户列表
	comments, err := s.repo.GetCommentsByMomentId(ctx, momentId)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] comment userIds err mid:%d", momentId)
	}
	comment := &model.MomentCommentModel{
		Uid:      model.Uid{UserId: userId},
		ReplyId:  replyId,
		MomentId: momentId,
		Content:  content,
	}
	_, err = s.repo.CommentCreate(ctx, comment)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] create comment err")
	}
	// 通知作者
	msg, err := app.NewMessagePack(message.EventMoment, &message.Moment{
		UserId: userId,
		Avatar: u.Avatar,
		Type:   "comment",
	})
	if err != nil {
		return err
	}
	userIds := []uint32{authorId}
	// 发送其他点赞好友
	for _, commentModel := range *comments {
		userIds = append(userIds, commentModel.UserId)
	}
	return s.PushUserIds(ctx, userIds, message.EventMoment, msg)
}

// momentCheck 验证动态id是否合法
func (s *Service) momentCheck(ctx context.Context, userId, momentId uint32) (user *model.UserModel, authorId uint32, err error) {
	user, err = s.repo.GetUserById(ctx, userId)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "[service.moment] user err id:%d", userId)
	}
	// 此条动态发布者
	moment, err := s.repo.GetMomentById(ctx, momentId)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "[service.moment] auther err mid:%d", momentId)
	}
	if moment.ID == 0 {
		return nil, 0, ErrMomentNotFound
	}
	if moment.SeeType != model.MomentSeeTypeAll { //非公开动态进一步判断权限
		// 是否存在或是否有权限
		exist, err := s.repo.TimelineExist(ctx, userId, momentId)
		if err != nil {
			return nil, 0, errors.Wrapf(err, "[service.moment] exist err uid:%d mid:%d", userId, momentId)
		}
		if !exist {
			return nil, 0, ErrMomentNotFound
		}
	}
	return user, moment.UserId, nil
}

// pushMessage 推送消息
func (s *Service) pushMessage(ctx context.Context, userId uint32, lines []*model.MomentTimelineModel, remind []uint32) (err error) {
	u, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] user err")
	}
	m := &message.Moment{
		UserId: userId,
		Avatar: u.Avatar,
		Type:   "new",
	}
	msg, err := app.NewMessagePack(message.EventMoment, m)
	if err != nil {
		return err
	}
	userIds := make([]uint32, len(lines))
	for _, line := range lines {
		userIds = append(userIds, line.UserId)
	}
	// 推送消息
	if err = s.PushUserIds(ctx, userIds, message.EventMoment, msg); err != nil {
		return errors.Wrapf(err, "[service.moment] push")
	}
	if len(remind) > 0 { // 是否需要提醒好友
		msg, err := app.NewMessagePack(message.EventMoment, &message.Moment{
			UserId: userId,
			Avatar: u.Avatar,
			Type:   "remind",
		})
		msg, err = app.NewMessagePack(message.EventMoment, m)
		if err != nil {
			return err
		}
		uids := make([]uint32, len(remind))
		for _, uid := range remind {
			uids = append(uids, uid)
		}
		// 推送消息
		err = s.PushUserIds(ctx, uids, message.EventMoment, msg)
		if err != nil {
			return errors.Wrapf(err, "[service.moment] push")
		}
	}
	return nil
}
