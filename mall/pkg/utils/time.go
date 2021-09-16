package utils

import (
	"fmt"
	"strconv"
	"time"
)

//TimeLayout 常用日期格式化模板
const (
	TimeLayout = "2006-01-02 15:04:05"
	DateSecLayout = "2006-01-02 15:04"
	DateLayout        = "2006-01-02"
)

// CurrentDate 获取当前时间
func CurrentDate() string {
	return time.Now().Format(TimeLayout)
}

//FormatToTimestamp 时间字符串转化为int64时间戳
func FormatToTimestamp(dateTime string) (int64, error) {
	loc, err := time.LoadLocation("Local") //获取时区
	if err != nil {
		return 0, err
	}
	tmp, err := time.ParseInLocation(TimeLayout, dateTime, loc)
	if err != nil {
		return 0, err
	}
	return tmp.Unix(), nil
}

// GetTodayDateInt 获取整形的日期
func GetTodayDateInt() int {
	dateStr := time.Now().Format("200601")
	date, err := strconv.Atoi(dateStr)
	if err != nil {
		return 0
	}
	return date
}

// TimeUnixFormat 格式化指定时间
func TimeUnixFormat(ux int64, layout string) string {
	return time.Unix(ux, 0).Format(layout)
}

// TimeToString 时间转字符串
func TimeToString(ts time.Time) string {
	return time.Unix(ts.Unix(), 00).Format(TimeLayout)
}

// TimeToShortString 时间转日期
func TimeToShortString(ts time.Time) string {
	return time.Unix(ts.Unix(), 00).Format(DateLayout)
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
