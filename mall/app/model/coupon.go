package model

const (
	//CouponStatusNormal 正常可领取
	CouponStatusNormal = 1 + iota
	//CouponStatusInvalid 已失效
	CouponStatusInvalid
	//CouponStatusReceived 已领取
	CouponStatusReceived
	//CouponStatusUsed 已使用
	CouponStatusUsed
)

const (
	//CouponTypeReduce 满减
	CouponTypeReduce = 1 + iota
	//CouponTypeDiscount 打折
	CouponTypeDiscount
)

//CouponModel 优惠券模型
type CouponModel struct {
	PriID
	Name     string `json:"name" gorm:"column:name;not null;type:varchar(50);comment:名称"`
	Type     int8   `json:"type" gorm:"column:type;not null;type:tinyint unsigned;comment:类型"`
	Value    int    `json:"value" gorm:"column:value;not null;type:int unsigned;comment:值"`
	Total    int    `json:"total" gorm:"column:total;not null;type:int unsigned;comment:发现总数"`
	MinPrice int    `json:"min_price" gorm:"column:min_price;not null;type:int unsigned;comment:最低价格"`
	StartAt  int    `json:"start_at" gorm:"column:start_at;not null;type:int unsigned;comment:开始时间"`
	EndAt    int    `json:"end_at" gorm:"column:end_at;not null;type:int unsigned;comment:结束时间"`
	Desc     string `json:"desc" gorm:"column:desc;not null;type:varchar(255);default:'';comment:描述"`
	Release
	OrderBy
	UpdateTime
}

// TableName 表名
func (u *CouponModel) TableName() string {
	return "coupon"
}

//Coupon 对外优惠券结构
type Coupon struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Type     int8    `json:"type"`
	Value    float64 `json:"value"`
	MinPrice float64 `json:"min_price"`
	StartAt  int     `json:"start_at"`
	EndAt    int     `json:"end_at"`
	Desc     string  `json:"desc"`
	Status   int     `json:"status"`
}
