package model

import "common/orm"

//SpuImageModel spu图片
type SpuImageModel struct {
	orm.PriID
	orm.Spu
	Name      string `json:"name" gorm:"column:name;not null;type:varchar(255);default:'';comment:图片名"`
	Img       string `json:"img" gorm:"column:img;not null;type:varchar(128);comment:图片地址"`
	IsDefault int8   `json:"is_default" gorm:"column:is_default;not null;default:0;comment:是否默认"`
	orm.OrderBy
}

// TableName 表名
func (u *SpuImageModel) TableName() string {
	return "pms_spu_image"
}