package model

import "encoding/json"


const (
	//AppPageHome 首页
	AppPageHome = 1 + iota
	//AppPageSearch 搜索页
	AppPageSearch
)

//AppSettingModel 应用设置模型
type AppSettingModel struct {
	PriID
	Page  int8   `json:"page" gorm:"column:page;not null;type:tinyint unsigned;comment:页面"`
	Type  int8   `json:"type" gorm:"column:type;not null;type:tinyint unsigned;comment:类型"`
	CatID int8   `json:"cat_id" gorm:"column:cat_id;not null;type:tinyint unsigned;comment:所属分类"`
	Data  string `json:"data" gorm:"column:data;not null;type:varchar(5000);default:'';comment:数据"`
	OrderBy
}

// TableName 表名
func (u *AppSettingModel) TableName() string {
	return "app_setting"
}

//AppSetting 对外页面配置数据结构
type AppSetting struct {
	Type int8            `json:"type"`
	Data json.RawMessage `json:"data"`
}
