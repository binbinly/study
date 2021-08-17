package model

//UserProfileModel 用户详情模型
type UserProfileModel struct {
	PriID
	UID
	Gender     int8   `gorm:"column:gender;not null;default:1;comment:性别" json:"gender"`
	Emotion    int8   `gorm:"column:emotion;not null;default:1;comment:情感" json:"emotion"`
	Job        string `gorm:"column:job;not null;type:varchar(60);default:'';comment:职业" json:"job"`
	Hometown   string `gorm:"column:hometown;not null;type:varchar(60);default:'';comment:故乡" json:"hometown"`
	Birthday   string `gorm:"column:birthday;not null;type:varchar(20);default:'';comment:生日" json:"birthday"`
	Sign       string `gorm:"column:sign;not null;type:varchar(255);default:'';comment:签名" json:"sign"`
	Background string `gorm:"column:background;not null;type:varchar(128);default:'';comment:主页图" json:"background"`
}

// TableName 表名
func (u *UserProfileModel) TableName() string {
	return "user_profile"
}
