package model

//GoodsCategoryModel 商品分类模型
type GoodsCategoryModel struct {
	PriID
	PID  int    `json:"pid" gorm:"column:pid;not null;type:int unsigned;comment:上级分类"`
	Name string `json:"name" gorm:"column:name;not null;type:varchar(60);comment:分类名"`
	Desc string `json:"desc" gorm:"column:desc;not null;type:varchar(255);default:'';comment:描述"`
	OrderBy
	UpdateTime
}

// TableName 表名
func (u *GoodsCategoryModel) TableName() string {
	return "goods_category"
}

//GoodsCategory 商品分类对外暴露结构
type GoodsCategory struct {
	ID   int    `json:"id"`
	PID  int    `json:"pid"`
	Name string `json:"name"`
}

//GoodsCategoryTree 属性结构
type GoodsCategoryTree struct {
	ID    int                  `json:"id"`
	PID   int                  `json:"pid"`
	Name  string               `json:"name"`
	Sort  int16                `json:"sort"`
	Child []*GoodsCategoryTree `json:"child"`
}

// 按照 sort 从大到小排序
type GoodsCategorySort []*GoodsCategoryTree

func (a GoodsCategorySort) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a GoodsCategorySort) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a GoodsCategorySort) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Sort < a[i].Sort
}
