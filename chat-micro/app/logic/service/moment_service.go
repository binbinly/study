package service

import (
	"context"

	"github.com/pkg/errors"

	"chat-micro/app/constvar"
	"chat-micro/app/logic/idl"
	"chat-micro/app/logic/model"
	"chat-micro/app/message"
	"chat-micro/internal/orm"
	"chat-micro/pkg/util"
)

//IMoment 朋友圈接口
type IMoment interface {
	//MomentPush 发布朋友圈
	MomentPush(ctx context.Context, userID uint32, content, image, video, location string,
		t, sType int8, remind, see []uint32) error
	//MomentTimeline 我的朋友圈
	MomentTimeline(ctx context.Context, userID uint32, offset int) (*model.Moment, error)
	//MomentList 好友朋友圈
	MomentList(ctx context.Context, myID, userID uint32, offset int) (*model.Moment, error)
	//MomentLike 点赞
	MomentLike(ctx context.Context, userID, momentID uint32) error
	//MomentComment 评论
	MomentComment(ctx context.Context, userID, replyID, momentID uint32, content string) error
}

// MomentPush 发布朋友圈
func (s *Service) MomentPush(ctx context.Context, userID uint32, content, image, video, location string,
	t, sType int8, remind, see []uint32) (err error) {
	// 我的好友列表
	friends, err := s.repo.GetFriendAll(ctx, userID)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] friends err")
	}
	// 好友id列表
	friendIds := make([]uint32, 0, len(friends))
	for _, f := range friends {
		friendIds = append(friendIds, f.FriendID)
	}
	// 过滤非好友元素
	newRemind := util.FilterSmallUInt32Slice(friendIds, func(v uint32) bool {
		return util.InuInt32Slice(v, remind)
	})
	m := &model.MomentModel{
		UID:      orm.UID{UserID: userID},
		Content:  content,
		Image:    image,
		Video:    video,
		Location: location,
		Remind:   util.SliceUInt32ToString2(newRemind),
		Type:     t,
		SeeType:  sType,
		See:      util.SliceUInt32ToString2(see),
	}
	// 开启事务
	db := orm.GetDB()
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
		UID:      orm.UID{UserID: userID},
		MomentID: id,
		IsOwn:    1,
	})
	for _, f := range friends {
		if sType == model.MomentSeeTypeNone {
			continue
		}
		line := &model.MomentTimelineModel{
			UID:      orm.UID{UserID: f.FriendID},
			MomentID: id,
			IsOwn:    0,
		}
		if sType == model.MomentSeeTypeAll {
			lines = append(lines, line)
		} else if sType == model.MomentSeeTypeOnly {
			if util.InuInt32Slice(f.FriendID, see) {
				lines = append(lines, line)
			}
		} else if sType == model.MomentSeeTypeExcept {
			if !util.InuInt32Slice(f.FriendID, see) {
				lines = append(lines, line)
			}
		}
	}
	if _, err = s.repo.TimelineBatchCreate(ctx, tx, lines); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.moment] timeline batch create err")
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.moment] tx commit err")
	}
	return s.pushMessage(ctx, userID, lines, newRemind)
}

// MomentTimeline 我的朋友圈动态
func (s *Service) MomentTimeline(ctx context.Context, userID uint32, offset int) (*model.Moment, error) {
	u, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	// 朋友圈动态
	mList, err := s.repo.GetMyMoments(ctx, userID, offset, constvar.DefaultLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] list err uid:%d", userID)
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
		mapUserIds[momentModel.UserID] = true
	}
	// 点赞信息
	likeList, err := s.repo.GetLikesByMomentIds(ctx, mIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] likes err mids:%v", mIds)
	}
	for _, like := range likeList {
		for _, uid := range like {
			mapUserIds[uid] = true
		}
	}
	// 评论信息
	commentList, err := s.repo.GetCommentsByMomentIds(ctx, mIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] comments err mids:%v", mIds)
	}
	for _, list := range commentList {
		for _, commentModel := range list {
			mapUserIds[commentModel.UserID] = true
			if commentModel.ReplyID > 0 {
				mapUserIds[commentModel.ReplyID] = true
			}
		}
	}
	// 汇总所有动态，点赞，评论，恢复的用户id
	userIds := make([]uint32, 0, len(mapUserIds))
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
func (s *Service) MomentList(ctx context.Context, myID, userID uint32, offset int) (*model.Moment, error) {
	u, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	// 朋友圈动态
	mList, err := s.repo.GetMomentsByUserID(ctx, myID, userID, offset, constvar.DefaultLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] list err uid:%d", userID)
	}
	if len(mList) == 0 {
		return &model.Moment{
			User: idl.TransferUserBase(u),
			List: make([]*model.MomentList, 0),
		}, nil
	}
	// 我的好友列表
	friends, err := s.repo.GetFriendAll(ctx, myID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] friends err")
	}
	// 好友id列表
	friendIds := make([]uint32, 0, len(friends))
	for _, f := range friends {
		friendIds = append(friendIds, f.FriendID)
	}
	mIds := make([]uint32, 0, len(mList))
	// 先用map存放，为去重用户id
	mapUserIds := make(map[uint32]bool, 0)
	for _, moment := range mList {
		mIds = append(mIds, moment.ID)
		mapUserIds[moment.UserID] = true
	}
	// 点赞信息
	likeList, err := s.repo.GetLikesByMomentIds(ctx, mIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] likes err mids:%v", mIds)
	}
	for _, like := range likeList {
		for _, uid := range like {
			if util.InuInt32Slice(uid, friendIds) { // 过滤非我的好友的点赞
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
		for _, comment := range list {
			if util.InuInt32Slice(comment.UserID, friendIds) { // 过滤非我的好友的评论
				mapUserIds[comment.UserID] = true
				if comment.ReplyID > 0 {
					if util.InuInt32Slice(comment.ReplyID, friendIds) {
						mapUserIds[comment.ReplyID] = true
					}
				}
			}
		}
	}
	// 汇总所有动态，点赞，评论，恢复的用户id
	userIds := make([]uint32, 0, len(mapUserIds))
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
func (s *Service) MomentLike(ctx context.Context, userID, momentID uint32) error {
	u, authorID, err := s.momentCheck(ctx, userID, momentID)
	if err != nil {
		return err
	}
	// 已经点赞的用户列表
	likeIds, err := s.repo.GetLikeUserIdsByMomentID(ctx, momentID)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] userIds err mid:%d", momentID)
	}
	// 是否已点赞
	isLike, err := s.repo.LikeExist(ctx, userID, momentID)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] exist like err uid:%d mid:%d", userID, momentID)
	}
	if isLike { // 已点赞，即取消
		if err = s.repo.LikeDelete(ctx, userID, momentID); err != nil {
			return errors.Wrapf(err, "[service.moment] delete err uid:%d mid:%d", userID, momentID)
		}
		return nil
	}
	// 创建点赞记录
	mLike := &model.MomentLikeModel{
		UserID:   userID,
		MomentID: momentID,
	}

	if _, err = s.repo.LikeCreate(ctx, mLike); err != nil {
		return errors.Wrapf(err, "[service.moment] create like err uid:%d mid:%d", userID, momentID)
	}
	// 通知作者
	userIds := []uint32{authorID}
	// 发送其他点赞好友
	for _, id := range likeIds {
		userIds = append(userIds, id)
	}
	return s.pushUserIds(ctx, userIds, message.EventMoment, &message.Moment{
		UserID: userID,
		Avatar: u.Avatar,
		Type:   "like",
	})
}

// MomentComment 评论
func (s *Service) MomentComment(ctx context.Context, userID, replyID, momentID uint32, content string) error {
	u, authorID, err := s.momentCheck(ctx, userID, momentID)
	if err != nil {
		return err
	}
	// 已评论用户列表
	comments, err := s.repo.GetCommentsByMomentID(ctx, momentID)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] comment userIds err mid:%d", momentID)
	}
	comment := &model.MomentCommentModel{
		UID:      orm.UID{UserID: userID},
		ReplyID:  replyID,
		MomentID: momentID,
		Content:  content,
	}
	_, err = s.repo.CommentCreate(ctx, comment)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] create comment err")
	}
	// 通知作者
	userIds := []uint32{authorID}
	// 发送其他点赞好友
	for _, commentModel := range comments {
		userIds = append(userIds, commentModel.UserID)
	}
	return s.pushUserIds(ctx, userIds, message.EventMoment, &message.Moment{
		UserID: userID,
		Avatar: u.Avatar,
		Type:   "comment",
	})
}

// momentCheck 验证动态id是否合法
func (s *Service) momentCheck(ctx context.Context, userID, momentID uint32) (user *model.UserModel, authorID uint32, err error) {
	user, err = s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "[service.moment] user err id:%d", userID)
	}
	// 此条动态发布者
	moment, err := s.repo.GetMomentByID(ctx, momentID)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "[service.moment] auther err mid:%d", momentID)
	}
	if moment.ID == 0 {
		return nil, 0, ErrMomentNotFound
	}
	if moment.SeeType != model.MomentSeeTypeAll { //非公开动态进一步判断权限
		// 是否存在或是否有权限
		exist, err := s.repo.TimelineExist(ctx, userID, momentID)
		if err != nil {
			return nil, 0, errors.Wrapf(err, "[service.moment] exist err uid:%d mid:%d", userID, momentID)
		}
		if !exist {
			return nil, 0, ErrMomentNotFound
		}
	}
	return user, moment.UserID, nil
}

// pushMessage 推送 朋友圈新动态消息
func (s *Service) pushMessage(ctx context.Context, userID uint32, lines []*model.MomentTimelineModel, remind []uint32) (err error) {
	u, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] user err")
	}
	userIds := make([]uint32, len(lines))
	for i, line := range lines {
		if line.UserID != userID { //不需要给自己发送 新动态通知
			userIds[i] = line.UserID
		}
	}
	// 推送消息
	if err = s.pushUserIds(ctx, userIds, message.EventMoment, &message.Moment{
		UserID: userID,
		Avatar: u.Avatar,
		Type:   "new",
	}); err != nil {
		return errors.Wrapf(err, "[service.moment] push")
	}
	if len(remind) > 0 { // 是否需要提醒好友
		uids := make([]uint32, 0, len(remind))
		for _, uid := range remind {
			uids = append(uids, uid)
		}
		// 推送消息
		if err = s.pushUserIds(ctx, uids, message.EventMoment, &message.Moment{
			UserID: userID,
			Avatar: u.Avatar,
			Type:   "remind",
		}); err != nil {
			return errors.Wrapf(err, "[service.moment] push")
		}
	}
	return nil
}
