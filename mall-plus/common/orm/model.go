package orm

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

	//ReleaseYes 已发布
	ReleaseYes = 1
	//DefaultOrder 默认排序
	DefaultOrder = "id DESC"
	//DefaultOrderSort 排序字段排序
	DefaultOrderSort = "sort DESC"
)

const (
	//StatusInit 状态-初始化
	StatusInit = iota
	//StatusSuccess 状态-成功
	StatusSuccess
	//StatusError 状态-失败
	StatusError
)

// UpdateTime 公共时间字段
type UpdateTime struct {
	CreatedAt int64 `gorm:"column:created_at;type:int(11) unsigned;not null;autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at;type:int(11) unsigned;not null;autoUpdateTime;comment:更新时间" json:"updated_at"`
}

// UpdateBy 公共操作者字段
type UpdateBy struct {
	CreateBy int `gorm:"column:create_by;type:int(11) unsigned;not null;default:0;comment:创建者" json:"create_by"`
	UpdateBy int `gorm:"column:update_by;type:int(11) unsigned;not null;default:0;comment:更新者" json:"update_by"`
}

// PriID 主键
type PriID struct {
	ID int64 `gorm:"primaryKey;autoIncrement;type:int(11) unsigned auto_increment;column:id;comment:ID" json:"id"`
}

// Create 创建时间
type Create struct {
	CreatedAt int64 `gorm:"column:created_at;type:int(11) unsigned;not null;autoCreateTime;comment:创建时间" json:"created_at"`
}

// Update 更新时间
type Update struct {
	UpdatedAt int64 `gorm:"column:updated_at;type:int(11) unsigned;not null;autoUpdateTime;comment:更新时间" json:"updated_at"`
}

// Delete 删除时间
type Delete struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`
}

// MID 用户ID
type MID struct {
	MemberID int64 `gorm:"column:member_id;not null;type:int(11) unsigned;index;comment:用户id" json:"member_id"`
}

//Spu 商品spu
type Spu struct {
	SpuID int64 `json:"spu_id" gorm:"column:spu_id;not null;index;type:int unsigned;comment:spu_id"`
}

//Sku 商品sku
type Sku struct {
	SkuID int64 `json:"sku_id" gorm:"column:sku_id;not null;index;type:int unsigned;comment:sku_id"`
}

//Cat 商品分类
type Cat struct {
	CatID int64 `json:"cat_id" gorm:"column:cat_id;not null;type:int unsigned;comment:产品分类"`
}

//OrderBy 排序字段
type OrderBy struct {
	Sort int32 `json:"sort" gorm:"column:sort;not null;type:int unsigned;default:50;comment:排序"`
}

//Release 是否发布
type Release struct {
	IsRelease int8 `json:"is_release" gorm:"column:is_release;not null;type:tinyint unsigned;default:0;comment:是否发布上线"`
}

//OffsetPage 分页查询
func OffsetPage(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

//WhereRelease 发布
func WhereRelease(db *gorm.DB) *gorm.DB {
	return db.Where("is_release=?", ReleaseYes)
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
