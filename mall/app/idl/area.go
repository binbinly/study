package idl

import (
	"mall/app/model"
)

//TransferArea 转换地区输出
func TransferArea(list []*model.Area) map[string]interface{} {
	province := make(map[int64]string, 0)
	city := make(map[int64]string, 0)
	county := make(map[int64]string, 0)
	for _, area := range list {
		switch area.Level {
		case 0:
			province[area.AreaCode] = area.Name
		case 1:
			city[area.AreaCode] = area.Name
		case 2:
			county[area.AreaCode] = area.Name
		}
	}
	return map[string]interface{}{
		"province": province,
		"city":     city,
		"county":   county,
	}
}
