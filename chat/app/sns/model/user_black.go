package model

//UserBlackModel 用户黑名单模型
type UserBlackModel struct {
	PriID
	UID
	BlackID uint32 `gorm:"column:black_id;not null;type:int(11) unsigned;index;comment:拉黑用户id" json:"black_id"`
}

// TableName 表名
func (u *UserBlackModel) TableName() string {
	return "user_black"
}
