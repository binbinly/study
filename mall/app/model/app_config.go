package model

const (
	//ConfigKeyHomeCat 首页分类键
	ConfigKeyHomeCat = "app_home_cat"
	//ConfigKeyPayList 支付配置
	ConfigKeyPayList = "app_pay_list"
)

//AppConfigModel 应用配置模型
type AppConfigModel struct {
	PriID
	Name  string `json:"name" gorm:"column:name;not null;type:char(15);comment:配置键"`
	Value string `json:"value" gorm:"column:value;not null;type:varchar(1000);comment:配置值"`
	Desc  string `json:"desc" gorm:"column:desc;not null;type:varchar(255);default:'';comment:描述"`
}

// TableName 表名
func (u *AppConfigModel) TableName() string {
	return "app_config"
}

//AppConfig 对外配置结构
type AppConfig struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//ConfigHomeCat 首页分类配置
type ConfigHomeCat struct {
	ID   int           `json:"id"`
	Name string        `json:"name"`
	List []*AppSetting `json:"list"`
}

//ConfigPayList 支付配置
type ConfigPayList struct {
	ID      int8   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
