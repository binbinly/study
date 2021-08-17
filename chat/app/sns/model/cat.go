package model

//CatModel 分类模型
type CatModel struct {
	PriID
	Name string `json:"name" gorm:"column:name;not null;type:varchar(120);comment:名称"`
	Sort int8   `json:"sort" gorm:"column:sort;not null;default:0;comment:排序"`
}

// TableName 表名
func (u *CatModel) TableName() string {
	return "cat"
}

//Cat 对外分类结构
type Cat struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
