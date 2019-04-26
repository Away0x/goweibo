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
		s = int(math.Floor(hour / 365))
		unit = "年"
	} else if hour > 30 {
		s = int(math.Floor(hour / 30))
		unit = "月"
	} else if hour > 0 {
		s = int(math.Floor(hour))
		unit = "小时"
	} else if minutes > 0 {
		s = int(math.Floor(minutes))
		unit = "分钟"
	} else if seconds > 0 {
		s = int(math.Floor(seconds))
		unit = "秒"
	}

	return strconv.Itoa(s) + unit + "前"
}
