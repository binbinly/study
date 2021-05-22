package proto

import "encoding/json"

const (
	EventLogin       = "login"
	EventSendCode    = "send_code"
	EventSearch      = "search"
	EventRegister    = "register"
	EventPhoneLogin  = "login_phone"
	EventUserEdit    = "user_edit"
	EventLogout      = "logout"
	EventUserReport  = "user_report"
	EventUserTags    = "user_tags"
	EventUserProfile = "user_profile"

	EventMomentCreate   = "moment_create"
	EventMomentComment  = "moment_comment"
	EventMomentLike     = "moment_like"
	EventMomentList     = "moment_list"
	EventMomentTimeline = "moment_timeline"

	EventGroupCreate       = "group_create"
	EventGroupInfo         = "group_info"
	EventGroupInvite       = "group_invite"
	EventGroupJoin         = "group_join"
	EventGroupKickoff      = "group_kickoff"
	EventGroupList         = "group_list"
	EventGroupQuit         = "group_quit"
	EventGroupEdit         = "group_edit"
	EventGroupEditNickname = "group_edit_nickname"
	EventGroupUser         = "group_user"

	EventFriendInfo       = "friend_info"
	EventFriendDestroy    = "friend_destroy"
	EventFriendList       = "friend_list"
	EventFriendTagList    = "friend_tag_list"
	EventFriendEditBlack  = "friend_edit_black"
	EventFriendEditStar   = "friend_edit_star"
	EventFriendEditAuth   = "friend_edit_auth"
	EventFriendEditRemark = "friend_edit_remark"

	EventCollectCreate  = "collect_create"
	EventCollectDestroy = "collect_destroy"
	EventCollectList    = "collect_list"

	EventChatDetail = "chat_detail"
	EventChatSend   = "chat_send"
	EventChatRecall = "chat_recall"

	EventApplyFriend = "apply_friend"
	EventApplyHandle = "apply_handle"
	EventApplyList   = "apply_list"
	EventApplyCount  = "apply_count"
)

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ReqRegister 注册
type ReqRegister struct {
	Phone    int64  `json:"phone"`    //手机号
	Username string `json:"username"` //用户名
	Password string `json:"password"` //密码
}

// ReqPhoneLogin 手机号登录
type ReqPhoneLogin struct {
	Phone      int64  `json:"phone"` // 手机号
	VerifyCode string `json:"code"`  // 验证码
}

// ReqUserEdit 修改用户信息
type ReqUserEdit struct {
	Avatar   string `json:"avatar"`   // 头像
	Nickname string `json:"nickname"` // 昵称
	Sign     string `json:"sign"`     // 签名
}

// ReqMomentCreate 发布朋友圈
type ReqMomentCreate struct {
	Content  string   `json:"content"`  // 内容
	Image    string   `json:"image"`    // 图片
	Video    string   `json:"video"`    // 视频
	Type     int8     `json:"type"`     // 类型 1=文本 2=图文 3=视频
	Location string   `json:"location"` // 地理位置
	Remind   []uint32 `json:"remind"`   // 提醒用户列表
	SeeType  int8     `json:"see_type"` // 可见类型 1=全部 2=私密 3=谁可见 4=谁不可见
	See      []uint32 `json:"see"`      // id列表
}

// ReqComment 评论
type ReqComment struct {
	Id      uint32 `json:"id"`       // 动态ID
	ReplyId uint32 `json:"reply_id"` // 回复者
	Content string `json:"content"`  // 内容
}

// ReqGroupEdit 修噶群组信息
type ReqGroupEdit struct {
	Id     uint32 `json:"id"`     // 群ID
	Name   string `json:"name"`   // 群名
	Remark string `json:"remark"` // 群公告
}

// ReqGroupNickname 修改群昵称
type ReqGroupNickname struct {
	Id       uint32 `json:"id"`       // 群ID
	Nickname string `json:"nickname"` // 群名
}

// ReqGroupAct 操作群用户
type ReqGroupAct struct {
	Id     uint32 `json:"id"`      // 群ID
	UserId uint32 `json:"user_id"` // 用户ID
}

// ReqFriendBlack 移入/移除黑名单
type ReqFriendBlack struct {
	UserId uint32 `json:"user_id"` //用户ID
	Black  int8   `json:"black"`   // 是否拉黑
}

// ReqFriendStar 设置/取消星标好友
type ReqFriendStar struct {
	UserId uint32 `json:"user_id"` //用户ID
	Star   int8   `json:"star"`    // 是否星标用户
}

// ReqFriendAuth 设置朋友圈权限
type ReqFriendAuth struct {
	UserId  uint32 `json:"user_id"`  //用户ID
	LookMe  int8   `json:"look_me"`  //看我
	LookHim int8   `json:"look_him"` //看他
}

// ReqFriendRemark 设置好友备注标签
type ReqFriendRemark struct {
	UserId   uint32   `json:"user_id"`  //用户ID
	Nickname string   `json:"nickname"` //备注内侧
	Tags     []string `json:"tags"`     //标签
}

// ReqReport 好友举报
type ReqReport struct {
	UserId   uint32 `json:"user_id"`  //用户ID
	Type     int8   `json:"type"`     // 1=用户，2=群组
	Content  string `json:"content"`  // 举报内容
	Category string `json:"category"` // 举报分类
}

// ReqCollectCreate 创建收藏
type ReqCollectCreate struct {
	Type    int8            `json:"type"`    // 聊天信息类型
	Content string          `json:"content"` // 内容
	Options json.RawMessage `json:"options"` // 额外选项
}

// ReqChatDetail 聊天详情
type ReqChatDetail struct {
	Id   uint32 `json:"id"`   //用户/群组ID
	Type int    `json:"type"` // 聊天类型，1=用户，2=群组
}

// ReqChatSend 发送消息
type ReqChatSend struct {
	ToId     uint32          `json:"to_id"`     // 用户/群组ID
	ChatType int             `json:"chat_type"` // 聊天类型，1=用户，2=群组
	Type     int             `json:"type"`      // 聊天信息类型
	Content  string          `json:"content"`   // 内容
	Options  json.RawMessage `json:"options"`   // 额外选项
}

// ReqChatRecall 撤回消息
type ReqChatRecall struct {
	Id       string `json:"id"`        // 消息id
	ToId     uint32 `json:"to_id"`     // 用户/群组ID
	ChatType int    `json:"chat_type"` // 聊天类型，1=用户，2=群组
}

// ReqApplyFriend 申请好友
type ReqApplyFriend struct {
	FriendId uint32 `json:"friend_id"` //好友ID
	Nickname string `json:"nickname"`  //备注昵称
	LookMe   int8   `json:"look_me"`   //看我
	LookHim  int8   `json:"look_him"`  //看他
}

// ReqApplyHandle 处理好友申请
type ReqApplyHandle struct {
	FriendId uint32 `json:"friend_id"` //好友ID
	Nickname string `json:"nickname"`  //备注内侧
	LookMe   int8   `json:"look_me"`   //看我
	LookHim  int8   `json:"look_him"`  //看他
}
