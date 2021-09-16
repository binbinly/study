package constvar

import (
	"fmt"
	"strconv"

	"mall/app/conf"
)

//ParseAmount 金额转元输出
func ParseAmount(amount int) float64 {
	f, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(amount)/100), 64)
	return f
}

//BuildResUrl 资源图片路径
func BuildResUrl(url string) string {
	return conf.Conf.App.DfsUrl + "/group1/" + url
}