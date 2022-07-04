package datetime

import (
	"fmt"
	"strconv"
	"time"
)

// GetTodayDateInt 获取整形的日期
func GetTodayDateInt() int {
	dateStr := time.Now().Format("20060102")
	date, err := strconv.Atoi(dateStr)
	if err != nil {
		return 0
	}
	return date
}

// TimeToString 时间转字符串
func TimeToString(ts time.Time) string {
	return time.Unix(ts.Unix(), 00).Format("2006-01-02 15:04:05")
}

// TimeToShortString 时间转日期
func TimeToShortString(ts time.Time) string {
	return time.Unix(ts.Unix(), 00).Format("2006.01.02")
}

// GetShowTime 格式化时间
func GetShowTime(ts time.Time) string {
	duration := time.Now().Unix() - ts.Unix()
	timeStr := ""
	if duration < 60 {
		timeStr = "刚刚发布"
	} else if duration < 3600 {
		timeStr = fmt.Sprintf("%d分钟前更新", duration/60)
	} else if duration < 86400 {
		timeStr = fmt.Sprintf("%d小时前更新", duration/3600)
	} else if duration < 86400*2 {
		timeStr = "昨天更新"
	} else {
		timeStr = TimeToShortString(ts) + "前更新"
	}
	return timeStr
}
