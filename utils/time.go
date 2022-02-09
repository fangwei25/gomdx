package utils

import (
	"time"
)

const (
	OneDaySeconds = 24 * 3600
)

var Location *time.Location

func init() {
	Location, _ = time.LoadLocation("Asia/Shanghai")
}

// IsSameDay 判断两个时间是否是同一天
func IsSameDay(t1 time.Time, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

// IsToday 判断时间是否和今天是同一天
func IsToday(t time.Time) bool {
	return IsSameDay(t, time.Now())
}

// IsTodayTimestamp 判断时间是否和今天是同一天
func IsTodayTimestamp(ts int64) bool {
	t := time.Unix(ts, 0)
	return IsSameDay(t, time.Now())
}

// GetTodayZeroTimestamp 获取今天0点的时间戳
func GetTodayZeroTimestamp() int64 {
	return GetDayZeroTimestamp(time.Now())
}

// GetDayZeroTimestamp 获取指定时间当天的0点时间戳
func GetDayZeroTimestamp(t time.Time) int64 {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

// TimeFormatPattern 全服统一的日期格式化
const TimeFormatPattern = "2006-01-02"
const TimeHourFormatPattern = "2006-01-02-15"

func GetDateStr() string {
	return GetDateStrByTime(time.Now())
}

// GetDateStrByTime 根据时间返回日期字符串
func GetDateStrByTime(t time.Time) string {
	return t.Format(TimeFormatPattern)
}

// GetDateStrByTimestamp 根据时间戳返回日期字符串
func GetDateStrByTimestamp(ts int64) string {
	t := time.Unix(ts, 0)
	return t.Format(TimeFormatPattern)
}

// GetDateHourStrByTime 根据时间返回日期小时字符串
func GetDateHourStrByTime(t time.Time) string {
	return t.Format(TimeHourFormatPattern)
}

const SecondsOfDay = 24 * 3600

// DiffNatureDays 计算两个日期间隔的自然天数
func DiffNatureDays(t1Time, t2Time time.Time) int {
	t1 := t1Time.Unix()
	t2 := t2Time.Unix()
	if t1 == t2 {
		return 0
	}
	if t1 > t2 {
		t1, t2 = t2, t1
	}

	diffDays := 0
	secDiff := t2 - t1
	if secDiff > SecondsOfDay {
		tmpDays := int(secDiff / SecondsOfDay)
		t1 += int64(tmpDays) * SecondsOfDay
		diffDays += tmpDays
	}

	st := time.Unix(t1, 0)
	et := time.Unix(t2, 0)
	dateFormatTpl := "20060102"
	if st.Format(dateFormatTpl) != et.Format(dateFormatTpl) {
		diffDays += 1
	}

	return diffDays
}
