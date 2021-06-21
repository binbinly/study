package proto

import "encoding/json"

const (
	//EventLogin 登录
	EventLogin       = "login"
	//EventSendCode 发送验证码
	EventSendCode    = "send_code"
	//EventSearch 搜索
	EventSearch      = "search"
	//EventRegister 注册
	EventRegister    = "register"
	//EventPhoneLogin 手机登录
	EventPhoneLogin  = "login_phone"
	//EventUserEdit 修改用户信息
	EventUserEdit    = "user_edit"
	//EventLogout 登出
	EventLogout      = "logout"
	//EventUserReport 举报
	EventUserReport  = "user_report"
	//EventUserTags 用户标签
	EventUserTags    = "user_tags"
	//EventUserProfile 用户详情
	EventUserProfile = "user_profile"

	//EventMomentCreate 发布朋友圈
	EventMomentCreate   = "moment_create"
	//EventMomentComment 朋友圈评论
	EventMomentComment  = "moment_comment"
	//EventMomentLike 朋友圈点赞
	EventMomentLike     = "moment_like"
	//EventMomentList 朋友圈列表
	EventMomentList     = "moment_list"
	//EventMomentTimeline 我的朋友圈时间线
	EventMomentTimeline = "moment_timeline"

	//EventGroupCreate 创建群组
	EventGroupCreate       = "group_create"
	//EventGroupInfo 群信息
	EventGroupInfo         = "group_info"
	//EventGroupInvite 邀请入群
	EventGroupInvite       = "group_invite"
	//EventGroupJoin 加入群
	EventGroupJoin         = "group_join"
	//EventGroupKickoff 踢出群
	EventGroupKickoff      = "group_kickoff"
	//EventGroupList 我的群列表
	EventGroupList         = "group_list"
	//EventGroupQuit 退出群
	EventGroupQuit         = "group_quit"
	//EventGroupEdit 修改群信息
	EventGroupEdit         = "group_edit"
	//EventGroupEditNickname 修改我在群中昵称
	EventGroupEditNickname = "group_edit_nickname"
	//EventGroupUser 群成员
	EventGroupUser         = "group_user"

	//EventFriendInfo 好友信息
	EventFriendInfo       = "friend_info"
	//EventFriendDestroy 删除好友
	EventFriendDestroy    = "friend_destroy"
	//EventFriendList 好友列表
	EventFriendList       = "friend_list"
	//EventFriendTagList 标签好友列表
	EventFriendTagList    = "friend_tag_list"
	//EventFriendEditBlack 黑名单加入/移出
	EventFriendEditBlack  = "friend_edit_black"
	//EventFriendEditStar 星标好友加入/移出
	EventFriendEditStar   = "friend_edit_star"
	//EventFriendEditAuth 设置好友朋友圈权限
	EventFriendEditAuth   = "friend_edit_auth"
	//EventFriendEditRemark 修改好友备注
	EventFriendEditRemark = "friend_edit_remark"

	//EventCollectCreate 添加收藏
	EventCollectCreate  = "collect_create"
	//EventCollectDestroy 删除收藏
	EventCollectDestroy = "collect_destroy"
	//EventCollectList 收藏列表
	EventCollectList    = "collect_list"

	//EventChatDetail 聊天回话详情
	EventChatDetail = "chat_detail"
	//EventChatSend 发送消息
	EventChatSend   = "chat_send"
	//EventChatRecall 聊天消息撤回
	EventChatRecall = "chat_recall"

	//EventApplyFriend 申请好友
	EventApplyFriend = "apply_friend"
	//EventApplyHandle 好友申请处理
	EventApplyHandle = "apply_handle"
	//EventApplyList 申请好友列表
	EventApplyList   = "apply_list"
	//EventApplyCount 未处理申请数
	EventApplyCount  = "apply_count"
)

// ReqLogin 登录
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
	ID      uint32 `json:"id"`       // 动态ID
	ReplyID uint32 `json:"reply_id"` // 回复者
	Content string `json:"content"`  // 内容
}

// ReqGroupEdit 修噶群组信息
type ReqGroupEdit struct {
	ID     uint32 `json:"id"`     // 群ID
	Name   string `json:"name"`   // 群名
	Remark string `json:"remark"` // 群公告
}

// ReqGroupNickname 修改群昵称
type ReqGroupNickname struct {
	ID       uint32 `json:"id"`       // 群ID
	Nickname string `json:"nickname"` // 群名
}

// ReqGroupAct 操作群用户
type ReqGroupAct struct {
	ID     uint32 `json:"id"`      // 群ID
	UserID uint32 `json:"user_id"` // 用户ID
}

// ReqFriendBlack 移入/移除黑名单
type ReqFriendBlack struct {
	UserID uint32 `json:"user_id"` //用户ID
	Black  int8   `json:"black"`   // 是否拉黑
}

// ReqFriendStar 设置/取消星标好友
type ReqFriendStar struct {
	UserID uint32 `json:"user_id"` //用户ID
	Star   int8   `json:"star"`    // 是否星标用户
}

// ReqFriendAuth 设置朋友圈权限
type ReqFriendAuth struct {
	UserID  uint32 `json:"user_id"`  //用户ID
	LookMe  int8   `json:"look_me"`  //看我
	LookHim int8   `json:"look_him"` //看他
}

// ReqFriendRemark 设置好友备注标签
type ReqFriendRemark struct {
	UserID   uint32   `json:"user_id"`  //用户ID
	Nickname string   `json:"nickname"` //备注内侧
	Tags     []string `json:"tags"`     //标签
}

// ReqReport 好友举报
type ReqReport struct {
	UserID   uint32 `json:"user_id"`  //用户ID
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
	ID   uint32 `json:"id"`   //用户/群组ID
	Type int    `json:"type"` // 聊天类型，1=用户，2=群组
}

// ReqChatSend 发送消息
type ReqChatSend struct {
	ToID     uint32          `json:"to_id"`     // 用户/群组ID
	ChatType int             `json:"chat_type"` // 聊天类型，1=用户，2=群组
	Type     int             `json:"type"`      // 聊天信息类型
	Content  string          `json:"content"`   // 内容
	Options  json.RawMessage `json:"options"`   // 额外选项
}

// ReqChatRecall 撤回消息
type ReqChatRecall struct {
	ID       string `json:"id"`        // 消息id
	ToID     uint32 `json:"to_id"`     // 用户/群组ID
	ChatType int    `json:"chat_type"` // 聊天类型，1=用户，2=群组
}

// ReqApplyFriend 申请好友
type ReqApplyFriend struct {
	FriendID uint32 `json:"friend_id"` //好友ID
	Nickname string `json:"nickname"`  //备注昵称
	LookMe   int8   `json:"look_me"`   //看我
	LookHim  int8   `json:"look_him"`  //看他
}

// ReqApplyHandle 处理好友申请
type ReqApplyHandle struct {
	FriendID uint32 `json:"friend_id"` //好友ID
	Nickname string `json:"nickname"`  //备注内侧
	LookMe   int8   `json:"look_me"`   //看我
	LookHim  int8   `json:"look_him"`  //看他
}
