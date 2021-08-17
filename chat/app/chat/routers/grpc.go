package routers

import (
	"chat/app/chat/handler/grpc/v1"
	"chat/app/chat/server"
	"chat/proto"
)

// NewGrpcRouter 实例化grpc路由
func NewGrpcRouter() *server.Engine {
	r := server.NewEngine()

	//用户
	r.AddRoute(proto.EventLogin, v1.Login)
	r.AddRoute(proto.EventRegister, v1.Register)
	r.AddRoute(proto.EventPhoneLogin, v1.LoginPhone)
	r.AddRoute(proto.EventLogout, v1.Logout)
	r.AddRoute(proto.EventUserProfile, v1.UserProfile)
	r.AddRoute(proto.EventSearch, v1.Search)
	r.AddRoute(proto.EventUserTags, v1.UserTags)
	r.AddRoute(proto.EventUserEdit, v1.UserEdit)
	r.AddRoute(proto.EventUserReport, v1.UserReport)

	//朋友圈
	r.AddRoute(proto.EventMomentCreate, v1.MomentCreate)
	r.AddRoute(proto.EventMomentComment, v1.MomentComment)
	r.AddRoute(proto.EventMomentLike, v1.MomentLike)
	r.AddRoute(proto.EventMomentList, v1.MomentList)
	r.AddRoute(proto.EventMomentTimeline, v1.MomentTimeline)

	//群组
	r.AddRoute(proto.EventGroupCreate, v1.GroupCreate)
	r.AddRoute(proto.EventGroupInfo, v1.GroupInfo)
	r.AddRoute(proto.EventGroupInvite, v1.GroupInvite)
	r.AddRoute(proto.EventGroupJoin, v1.GroupJoin)
	r.AddRoute(proto.EventGroupKickoff, v1.GroupKickoff)
	r.AddRoute(proto.EventGroupList, v1.GroupList)
	r.AddRoute(proto.EventGroupQuit, v1.GroupQuit)
	r.AddRoute(proto.EventGroupEdit, v1.GroupEdit)
	r.AddRoute(proto.EventGroupEditNickname, v1.GroupEditNickname)
	r.AddRoute(proto.EventGroupUser, v1.GroupUser)

	//好友
	r.AddRoute(proto.EventFriendInfo, v1.FriendInfo)
	r.AddRoute(proto.EventFriendDestroy, v1.FriendDestroy)
	r.AddRoute(proto.EventFriendList, v1.FriendList)
	r.AddRoute(proto.EventFriendTagList, v1.FriendTagList)
	r.AddRoute(proto.EventFriendEditBlack, v1.FriendEditBlack)
	r.AddRoute(proto.EventFriendEditStar, v1.FriendEditStar)
	r.AddRoute(proto.EventFriendEditAuth, v1.FriendEditAuth)
	r.AddRoute(proto.EventFriendEditRemark, v1.FriendEditRemark)

	//收藏
	r.AddRoute(proto.EventCollectCreate, v1.CollectCreate)
	r.AddRoute(proto.EventCollectDestroy, v1.CollectDestroy)
	r.AddRoute(proto.EventCollectList, v1.CollectList)

	//聊天会话
	r.AddRoute(proto.EventChatDetail, v1.ChatDetail)
	r.AddRoute(proto.EventChatSend, v1.ChatSend)
	r.AddRoute(proto.EventChatRecall, v1.ChatRecall)

	//好友申请
	r.AddRoute(proto.EventApplyFriend, v1.ApplyFriend)
	r.AddRoute(proto.EventApplyHandle, v1.ApplyHandle)
	r.AddRoute(proto.EventApplyList, v1.ApplyList)
	r.AddRoute(proto.EventApplyCount, v1.ApplyCount)

	return r
}
