package routers

import (
	"chat/app/logic/grpc"
	grpc2 "chat/app/logic/handler/grpc"
	"chat/proto"
)

// Grpc 路由定义
func NewRouter() *grpc.Engine {
	r := grpc.NewEngine()

	//用户
	r.AddRoute(proto.EventLogin, grpc2.Login)
	r.AddRoute(proto.EventRegister, grpc2.Register)
	r.AddRoute(proto.EventPhoneLogin, grpc2.LoginPhone)
	r.AddRoute(proto.EventLogout, grpc2.Logout)
	r.AddRoute(proto.EventSendCode, grpc2.SendCode)
	r.AddRoute(proto.EventUserProfile, grpc2.UserProfile)
	r.AddRoute(proto.EventSearch, grpc2.Search)
	r.AddRoute(proto.EventUserTags, grpc2.UserTags)
	r.AddRoute(proto.EventUserEdit, grpc2.UserEdit)
	r.AddRoute(proto.EventUserReport, grpc2.UserReport)

	//朋友圈
	r.AddRoute(proto.EventMomentCreate, grpc2.MomentCreate)
	r.AddRoute(proto.EventMomentComment, grpc2.MomentComment)
	r.AddRoute(proto.EventMomentLike, grpc2.MomentLike)
	r.AddRoute(proto.EventMomentList, grpc2.MomentList)
	r.AddRoute(proto.EventMomentTimeline, grpc2.MomentTimeline)

	//群组
	r.AddRoute(proto.EventGroupCreate, grpc2.GroupCreate)
	r.AddRoute(proto.EventGroupInfo, grpc2.GroupInfo)
	r.AddRoute(proto.EventGroupInvite, grpc2.GroupInvite)
	r.AddRoute(proto.EventGroupJoin, grpc2.GroupJoin)
	r.AddRoute(proto.EventGroupKickoff, grpc2.GroupKickoff)
	r.AddRoute(proto.EventGroupList, grpc2.GroupList)
	r.AddRoute(proto.EventGroupQuit, grpc2.GroupQuit)
	r.AddRoute(proto.EventGroupEdit, grpc2.GroupEdit)
	r.AddRoute(proto.EventGroupEditNickname, grpc2.GroupEditNickname)
	r.AddRoute(proto.EventGroupUser, grpc2.GroupUser)

	//好友
	r.AddRoute(proto.EventFriendInfo, grpc2.FriendInfo)
	r.AddRoute(proto.EventFriendDestroy, grpc2.FriendDestroy)
	r.AddRoute(proto.EventFriendList, grpc2.FriendList)
	r.AddRoute(proto.EventFriendTagList, grpc2.FriendTagList)
	r.AddRoute(proto.EventFriendEditBlack, grpc2.FriendEditBlack)
	r.AddRoute(proto.EventFriendEditStar, grpc2.FriendEditStar)
	r.AddRoute(proto.EventFriendEditAuth, grpc2.FriendEditAuth)
	r.AddRoute(proto.EventFriendEditRemark, grpc2.FriendEditRemark)

	//收藏
	r.AddRoute(proto.EventCollectCreate, grpc2.CollectCreate)
	r.AddRoute(proto.EventCollectDestroy, grpc2.CollectDestroy)
	r.AddRoute(proto.EventCollectList, grpc2.CollectList)

	//聊天会话
	r.AddRoute(proto.EventChatDetail, grpc2.ChatDetail)
	r.AddRoute(proto.EventChatSend, grpc2.ChatSend)
	r.AddRoute(proto.EventChatRecall, grpc2.ChatRecall)

	//好友申请
	r.AddRoute(proto.EventApplyFriend, grpc2.ApplyFriend)
	r.AddRoute(proto.EventApplyHandle, grpc2.ApplyHandle)
	r.AddRoute(proto.EventApplyList, grpc2.ApplyList)
	r.AddRoute(proto.EventApplyCount, grpc2.ApplyCount)

	return r
}
