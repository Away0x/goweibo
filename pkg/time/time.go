package time

import (
	"math"
	"strconv"
	"time"
)

// SinceForHuman 1小时前 -> 这样的展示方式
func SinceForHuman(t time.Time) string {
	duration := time.Since(t)
	hour := duration.Hours()
	minutes := duration.Minutes()
	seconds := duration.Seconds()

	unit := "秒"
	s := 0
	if hour > (365 * 24) {
		s = int(math.Floor(hour / (365 * 24)))
		unit = "年"
	} else if hour > (30 * 24) {
		s = int(math.Floor(hour / (30 * 24)))
		unit = "月"
	} else if hour > 24 {
		s = int(math.Floor(hour / 24))
		unit = "天"
	} else if hour > 1 {
		s = int(math.Floor(hour))
		unit = "小时"
	} else if minutes > 1 {
		s = int(math.Floor(minutes))
		unit = "分钟"
	} else if seconds > 0 {
		return "刚刚"
	}

	return strconv.Itoa(s) + unit + "前"
}
