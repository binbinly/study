package apply

// FriendParams 申请好友
type FriendParams struct {
	FriendID uint32 `json:"friend_id" binding:"required,numeric" example:"1"`         //好友ID
	Nickname string `json:"nickname"  binding:"required,min=1,max=30" example:"test"` //备注昵称
	LookMe   int8   `json:"look_me"  binding:"required,oneof=0 1" example:"1"`        //看我
	LookHim  int8   `json:"look_him" binding:"required,oneof=0 1" example:"1"`        //看他
}

// HandleParams 处理好友申请
type HandleParams struct {
	FriendID uint32 `json:"friend_id" binding:"required,numeric" example:"1"`         //好友ID
	Nickname string `json:"nickname"  binding:"required,min=1,max=30" example:"test"` //备注内侧
	LookMe   int8   `json:"look_me"  binding:"required,oneof=0 1" example:"1"`        //看我
	LookHim  int8   `json:"look_him" binding:"required,oneof=0 1" example:"1"`        //看他
}
