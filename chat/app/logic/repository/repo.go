package repository

import (
	"context"

	"gorm.io/gorm"

	"chat/app/logic/cache"
)

var _ IRepo = (*Repo)(nil)

//IRepo 数据仓库接口
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

// Repo mysql struct
type Repo struct {
	db *gorm.DB

	userCache    *cache.UserCache
	collectCache *cache.CollectCache
	tagCache     *cache.TagCache

	momentCache   *cache.MomentCache
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

		momentCache:   cache.NewMomentCache(),
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
