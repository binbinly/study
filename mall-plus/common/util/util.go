package util

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"common/conf"
	"common/constvar"
)

var dfsEndpoint string

func init() {
	dfsEndpoint = os.Getenv("MICRO_DFS_ENDPOINT")
}

//ParseAmount 金额转元输出
func ParseAmount(amount int) float64 {
	f, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(amount)/100), 64)
	return f
}

//FormatAmount 格式化金额为分
func FormatAmount(amount float64) int {
	return int(amount) * 100
}

//GetPageOffset 获取分页起始偏移量
func GetPageOffset(page int32) int {
	var offset int32
	if page > 0 {
		offset = (page - 1) * constvar.DefaultLimit
	}
	return int(offset)
}

//BuildResUrl 构建资源图片完整路径
func BuildResUrl(c conf.DFSConfig, url string) string {
	if dfsEndpoint != "" {
		c.Endpoint = dfsEndpoint
	}
	return strings.Join([]string{c.Endpoint, c.Bucket, url}, "/")
}

//BuildOrderNo 构建订单号
func BuildOrderNo() string {
	now := time.Now()
	orderNo := fmt.Sprintf("%04d%02d%02d%02d%02d%02d%08d", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), rand.Int31n(100000000))
	return orderNo
}
