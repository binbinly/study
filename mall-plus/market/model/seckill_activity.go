package model

import "common/orm"

//SeckillActivityModel 秒杀活动
type SeckillActivityModel struct {
	orm.PriID
	Title   string `json:"title" gorm:"column:title;not null;type:varchar(128);comment:标题"`
	Cover   string `json:"cover" gorm:"column:cover;not null;type:varchar(128);default:'';comment:活动封面"`
	StartAt int64  `json:"start_at" gorm:"column:start_at;not null;type:int;comment:开始时间"`
	EndAt   int64  `json:"end_at" gorm:"column:end_at;not null;type:int;comment:结束时间"`
	orm.Release
	orm.UpdateBy
	orm.UpdateTime
}

// TableName 表名
func (u *SeckillActivityModel) TableName() string {
	return "sms_seckill_activity"
}
