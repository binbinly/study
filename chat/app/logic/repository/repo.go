package repository

import (
	"chat/app/logic/cache"
	"context"

	"gorm.io/gorm"
)

var _ IRepo = (*Repo)(nil)

type IRepo interface {
	IApply
	ICollect
	IFriend
	IGroup
	IGroupUser
	IMessage
	IMoment
	IMomentComment
	IMomentTimeline
	IMomentLike
	IReport
	IUser
	IUserTag
	IReport
}

// repo mysql struct
type Repo struct {
	db *gorm.DB

	userCache    *cache.UserCache
	collectCache *cache.CollectCache
	tagCache     *cache.TagCache

	likeCache     *cache.LikeCache
	commentCache  *cache.CommentCache
	timelineCache *cache.TimelineCache

	applyCache  *cache.ApplyCache
	friendCache *cache.FriendCache

	groupCache     *cache.GroupCache
	groupUserCache *cache.GroupUserCache
}

// New new a Dao and return
func New(db *gorm.DB) IRepo {
	return &Repo{
		db:           db,
		userCache:    cache.NewUserCache(),
		collectCache: cache.NewCollectCache(),
		tagCache:     cache.NewTagCache(),

		likeCache:     cache.NewLikeCache(),
		commentCache:  cache.NewCommentCache(),
		timelineCache: cache.NewTimelineCache(),

		applyCache:  cache.NewApplyCache(),
		friendCache: cache.NewFriendCache(),

		groupCache:     cache.NewGroupCache(),
		groupUserCache: cache.NewGroupUserCache(),
	}
}

// Ping ping mysql
func (r *Repo) Ping(c context.Context) error {
	return nil
}

// Close release mysql connection
func (r *Repo) Close() error {
	return nil
}
