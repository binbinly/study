package model

import "common/orm"

//SpuDescModel spu介绍
type SpuDescModel struct {
	orm.PriID
	orm.Spu
	Content string `json:"content" gorm:"column:content;comment:商品介绍"`
}

// TableName 表名
func (u *SpuDescModel) TableName() string {
	return "pms_spu_desc"
}
