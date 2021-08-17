package repository

import (
	"gorm.io/gorm"

	"chat/app/chat/cache"
	"chat/internal/orm"
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
	IUserTag
	IReport
	IEmoticon

	Close() error
}

// Repo mysql struct
type Repo struct {
	db *gorm.DB

	collectCache *cache.CollectCache
	tagCache     *cache.TagCache
	emoCache     *cache.EmoticonCache

	momentCache   *cache.MomentCache
	likeCache     *cache.LikeCache
	commentCache  *cache.CommentCache
	timelineCache *cache.TimelineCache

	applyCache     *cache.ApplyCache
	friendCache    *cache.FriendCache
	friendAllCache *cache.FriendAllCache

	groupCache        *cache.GroupCache
	groupAllCache     *cache.GroupAllCache
	groupUserCache    *cache.GroupUserCache
	groupUserAllCache *cache.GroupUserAllCache
}

// New new a Dao and return
func New(db *gorm.DB) IRepo {
	return &Repo{
		db:           db,
		collectCache: cache.NewCollectCache(),
		tagCache:     cache.NewTagCache(),
		emoCache:     cache.NewEmoticonCache(),

		momentCache:   cache.NewMomentCache(),
		likeCache:     cache.NewLikeCache(),
		commentCache:  cache.NewCommentCache(),
		timelineCache: cache.NewTimelineCache(),

		applyCache:     cache.NewApplyCache(),
		friendCache:    cache.NewFriendCache(),
		friendAllCache: cache.NewFriendAllCache(),

		groupCache:        cache.NewGroupCache(),
		groupAllCache:     cache.NewGroupAllCache(),
		groupUserCache:    cache.NewGroupUserCache(),
		groupUserAllCache: cache.NewGroupUserAllCache(),
	}
}

// Close release mysql connection
func (r *Repo) Close() error {
	return orm.CloseDB()
}
