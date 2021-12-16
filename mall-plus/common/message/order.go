package message

//OrderMessage 订单消息
type OrderMessage struct {
	OrderID  int64 `json:"order_id"`  //订单id
	MemberID int64 `json:"member_id"` //会员id
}
