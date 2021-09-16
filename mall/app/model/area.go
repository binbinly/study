package model

//数据来源： https://github.com/kakuilan/china_area_mysql
//AreaModel 省市区
type AreaModel struct {
	PriID
	Level      int8    `json:"level" gorm:"column:level;not null;default:0;comment:层级"`
	AreaCode   int64   `json:"area_code" gorm:"column:area_code;not null;comment:行政代码"`
	ParentCode int64   `json:"parent_code" gorm:"column:parent_code;not null;comment:父级行政代码"`
	Shortname  string  `json:"shortname" gorm:"column:shortname;not null;varchar(20);comment:简称"`
	Name       string  `json:"name" gorm:"column:name;not null;type:varchar(50);comment:名称"`
	MergerName string  `json:"merger_name" gorm:"column:merger_name;not null;type:varchar(100);comment:全称"`
	Pinyin     string  `json:"pinyin" gorm:"column:pinyin;not null;type:varchar(30);comment:拼音"`
	CityCode   string  `json:"code" gorm:"column:code;not null;type:char(6);comment:区号"`
	ZipCode    int16   `json:"zip_code" gorm:"column:zip_code;not null;type:mediumint unsigned;comment:邮编"`
	Lng        float64 `json:"lng" gorm:"column:lng;not null;type:decimal(10,6);comment:经度"`
	Lat        float64 `json:"lat" gorm:"column:lat;not null;type:decimal(10,6);comment:纬度"`
}

// TableName 表名
func (u *AreaModel) TableName() string {
	return "area"
}

//Area 对外地区结构
type Area struct {
	Level    int8   `json:"level"`
	AreaCode int64  `json:"code"`
	Name     string `json:"name"`
}
