package model

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// NullType 空字节类型
type NullType byte

const (
	_ NullType = iota
	// IsNull the same as `is null`
	IsNull
	// IsNotNull the same as `is not null`
	IsNotNull

	ReleaseYes   = 1         //已发布
	DefaultOrder = "id DESC" //默认排序

	StatusInit    = 0 //状态-初始化
	StatusSuccess = 1 //状态-成功
	StatusError   = 2 //状态-失败
)

// UpdateTime 公共时间字段
type UpdateTime struct {
	CreatedAt int64 `gorm:"column:created_at;type:int(11) unsigned;not null;autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at;type:int(11) unsigned;not null;autoUpdateTime;comment:更新时间" json:"updated_at"`
}

// ID 主键
type PriID struct {
	ID uint32 `gorm:"primaryKey;autoIncrement;type:int(11) unsigned auto_increment;column:id;comment:ID" json:"id"`
}

// Create 创建时间
type Create struct {
	CreatedAt int64 `gorm:"column:created_at;type:int(11) unsigned;not null;autoCreateTime;comment:创建时间" json:"created_at"`
}

// Update 更新时间
type Update struct {
	UpdatedAt int64 `gorm:"column:updated_at;type:int(11) unsigned;not null;autoUpdateTime;comment:更新时间" json:"updated_at"`
}

// Uid 用户ID
type Uid struct {
	UserId uint32 `gorm:"column:user_id;not null;type:int(11) unsigned;index;comment:用户id" json:"user_id"`
}

//分页查询
func OffsetPage(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

// WhereBuild sql build where
// see: https://github.com/go-gorm/gorm/issues/2055
func WhereBuild(where map[string]interface{}) (whereSQL string, vals []interface{}, err error) {
	for k, v := range where {
		ks := strings.Split(k, " ")
		if len(ks) > 2 {
			return "", nil, fmt.Errorf("Error in query condition: %s. ", k)
		}

		if whereSQL != "" {
			whereSQL += " AND "
		}

		fmt.Println(strings.Join(ks, ","))
		switch len(ks) {
		case 1:
			//fmt.Println(reflect.TypeOf(v))
			switch v := v.(type) {
			case NullType:
				fmt.Println()
				if v == IsNotNull {
					whereSQL += fmt.Sprint(k, " IS NOT NULL")
				} else {
					whereSQL += fmt.Sprint(k, " IS NULL")
				}
			default:
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			}
		case 2:
			k = ks[0]
			switch ks[1] {
			case "=":
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			case ">":
				whereSQL += fmt.Sprint(k, ">?")
				vals = append(vals, v)
			case ">=":
				whereSQL += fmt.Sprint(k, ">=?")
				vals = append(vals, v)
			case "<":
				whereSQL += fmt.Sprint(k, "<?")
				vals = append(vals, v)
			case "<=":
				whereSQL += fmt.Sprint(k, "<=?")
				vals = append(vals, v)
			case "!=":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
			case "<>":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
			case "in":
				whereSQL += fmt.Sprint(k, " in (?)")
				vals = append(vals, v)
			case "like":
				whereSQL += fmt.Sprint(k, " like ?")
				vals = append(vals, v)
			}
		}
	}
	return
}
