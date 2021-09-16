package model

//CouponUserModel 用户优惠券模型
type CouponUserModel struct {
	PriID
	UID
	CouponID int  `json:"coupon_id" gorm:"column:coupon_id;not null;type:int unsigned;comment:优惠券id"`
	IsUsed   int8 `json:"is_used" gorm:"column:is_used;not null;type:tinyint unsigned;default:0;comment:是否使用"`
	UpdateTime
}

// TableName 表名
func (u *CouponUserModel) TableName() string {
	return "coupon_user"
}
