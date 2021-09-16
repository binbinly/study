package model

//UserLevelModel 用户等级
type UserLevelModel struct {
	PriID
	Name     string `json:"name" gorm:"column:name;not null;type:char(15);comment:配置键"`
	Level    int8   `json:"level" gorm:"column:level;not null;type:tinyint unsigned;comment:等级"`
	Discount int8   `json:"discount" gorm:"column:discount;not null;type:tinyint unsigned;comment:折扣比例*100"`
	MinPrice int    `json:"min_price" gorm:"column:min_price;not null;type:int unsigned;comment:消费最低金额"`
	MinCount int    `json:"min_count" gorm:"column:min_count;not null;type:int unsigned;comment:消费最低次数"`
}

// TableName 表名
func (u *UserLevelModel) TableName() string {
	return "user_level"
}