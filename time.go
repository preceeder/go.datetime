package datetime

import (
	"time"
)

type DateTime struct {
	time         time.Time
	weekStartsAt time.Weekday
}

func Now(zone ...string) *DateTime {
	return &DateTime{time: time.Now().In(timeZoneHandler(zone...)), weekStartsAt: time.Monday}
}

func (t *DateTime) Now(zone ...string) *DateTime {
	return &DateTime{time: time.Now().In(timeZoneHandler(zone...))}
}

func (t *DateTime) SetWeekStartsAt(weekday time.Weekday) *DateTime {
	t.weekStartsAt = weekday
	return t
}

func (t *DateTime) Copy() *DateTime {
	return &DateTime{time: t.time, weekStartsAt: t.weekStartsAt}
}

// 输出 2006-01-02 15:04:05
func (t *DateTime) String() string {
	return t.time.Format(time.DateTime)
}

func (t *DateTime) Format(layout string) string {
	return t.time.Format(layout)
}

func (t *DateTime) Timestamp() int64 {
	return t.time.Unix()
}
func (t *DateTime) Unix() int64 {
	return t.time.Unix()
}

func (t *DateTime) UnixMilli() int64 {
	return t.time.UnixMilli()
}

func (t *DateTime) UnixMicro() int64 {
	return t.time.UnixMicro()
}

func (t *DateTime) UnixNano() int64 {
	return t.time.UnixNano()
}

func (t *DateTime) AddDays(d int) *DateTime {
	t.time = t.time.AddDate(0, 0, d)
	return t
}

func (t *DateTime) SubDays(d int) *DateTime {
	t.time = t.time.AddDate(0, 0, -d)
	return t
}

func (t *DateTime) AddSeconds(seconds int) *DateTime {
	t.time = t.time.Add(time.Second * time.Duration(seconds))
	return t
}

func (t *DateTime) SubSeconds(seconds int) *DateTime {
	t.time = t.time.Add(-time.Second * time.Duration(seconds))
	return t
}

func (t *DateTime) AddMinutes(minutes int) *DateTime {
	t.time = t.time.Add(time.Minute * time.Duration(minutes))
	return t
}

func (t *DateTime) SubMinutes(minutes int) *DateTime {
	t.time = t.time.Add(-time.Minute * time.Duration(minutes))
	return t
}

func (t *DateTime) AddHours(hours int) *DateTime {
	t.time = t.time.Add(time.Hour * time.Duration(hours))
	return t
}
func (t *DateTime) SubHours(hours int) *DateTime {
	t.time = t.time.Add(-time.Hour * time.Duration(hours))
	return t
}

func (t *DateTime) WeekEndsAt() time.Weekday {
	return time.Weekday((int(t.weekStartsAt) + DaysPerWeek - 1) % DaysPerWeek)
}

func (t *DateTime) WeekStartsAt() time.Weekday {
	return t.weekStartsAt
}

// 本周的第几天
func (t *DateTime) DayOfWeek() int {
	return (int(t.time.Weekday())+DaysPerWeek-int(t.weekStartsAt))%DaysPerWeek + 1
}

// 本月的第几天
func (t *DateTime) DayOfMouth() int {
	return t.time.Day()
}

func (t *DateTime) DayOfYesterday() int {
	return t.time.YearDay()
}

func (t *DateTime) IsLeapYear() bool {
	year := t.time.Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// 今天剩余的时间 秒
func (t *DateTime) DayNextSecond() int64 {
	last := int64(t.EndOfDay().Sub(t.time) / time.Second)
	if last < 1 {
		last = 1
	}
	return last
}

// 今天开始后的时间 秒
func (t *DateTime) DayLastSecond() int64 {
	return int64(t.time.Sub(t.StartOfDay()) / time.Second)
}

// 本周剩余的时间   周一为开始时间， 周末为结束时间
func (t *DateTime) WeekNextSecond() int64 {
	last := int64(t.EndOfWeek().Sub(t.time) / time.Second)
	if last <= 1 {
		return 1
	}
	return last
}

// 获取本月的剩余时间
func (t *DateTime) MouthNextSecond() int64 {
	last := int64(t.EndOfMonth().Sub(t.time) / time.Second)
	if last <= 1 {
		return 1
	}
	return last
}

// 获取当前小时的剩余时间
func (t *DateTime) HourNextSecond() int64 {
	last := int64(t.EndOfHour().Sub(t.time) / time.Second)
	if last <= 1 {
		return 1
	}
	return last
}

// 获取当前分钟的剩余时间
func (t *DateTime) MinuteNextSecond() int64 {
	last := int64(t.EndOfMinute().Sub(t.time) / time.Second)
	if last <= 1 {
		return 1
	}
	return last
}

// 时间戳转化为时间类型
func TimestampToDateTime(timestamp int64, zone ...string) *DateTime {
	return &DateTime{
		time:         time.Unix(timestamp, MinNanosecond).In(timeZoneHandler(zone...)),
		weekStartsAt: time.Monday,
	}
}

// 时间string化为时间类型
func StrToDateTime(datetime string, format string, zone ...string) (*DateTime, error) {
	loc := timeZoneHandler(zone...)
	t, err := time.ParseInLocation(format, datetime, loc)
	return &DateTime{
		time:         t,
		weekStartsAt: time.Monday,
	}, err
}
