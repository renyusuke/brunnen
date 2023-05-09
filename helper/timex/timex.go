package timex

import (
	"fmt"
	"github.com/golang-module/carbon/v2"

	"time"
)

// GetMonthBetweenByStr 获取指定时间月份区间
func GetMonthBetweenByStr(date string) (int64, int64) {
	start := carbon.Parse(date).StartOfMonth().Millisecond()
	end := carbon.Parse(date).EndOfMonth().Millisecond()
	return int64(start), int64(end)
}

// GetMonthBetween 获取当前时间月份区间
func GetMonthBetween() (int64, int64) {
	start := carbon.Now().StartOfMonth().Millisecond()
	end := carbon.Now().EndOfMonth().Millisecond()
	return int64(start), int64(end)
}

// TimeStampFormatString 時間戳format 字串
func TimeStampFormatString(timeStr int64, formatKey string) (res string) {
	tm := time.Unix(timeStr, 0).Format(formatKey)
	res = fmt.Sprintf("%v", tm)
	return res
}
