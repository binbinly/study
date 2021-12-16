package model

import "common/orm"

//CollectSpuModel 用户收藏商品
type CollectSpuModel struct {
	orm.PriID
	orm.MID
	orm.Spu
	SpuName string `json:"spu_name" gorm:"column:spu_name;not null;type:varchar(255);comment:spu_name"`
	SpuImg  string `json:"spu_img" gorm:"column:spu_img;not null;type:varchar(255);comment:spu_img"`
	orm.Create
	orm.Delete
}

// TableName 表名
func (u *CollectSpuModel) TableName() string {
	return "ums_collect_spu"
}