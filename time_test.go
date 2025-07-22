package datetime

import (
	"fmt"
	"testing"
)

func TestDateTime_AddHours(t1 *testing.T) {
	d := Now()
	fmt.Println(d)
	fmt.Println(d.AddHours(1))

	fmt.Println(d.EndOfHour())
	fmt.Println(d.EndOfDay())
	fmt.Println(d.EndOfWeek())
	fmt.Println(d.EndOfMonth())
	fmt.Println(d.EndOfYear())
	fmt.Println(d.EndOfMinute())

	fmt.Println(d.StartOfMinute())
	fmt.Println(d.StartOfHour())
	fmt.Println(d.StartOfDay())
	fmt.Println(d.StartOfWeek())
	fmt.Println(d.StartOfMonth())
	fmt.Println(d.StartOfYear())

	fmt.Println(d.MinuteNextSecond())
	fmt.Println(d.HourNextSecond())
	fmt.Println(d.DayNextSecond())
	fmt.Println(d.WeekNextSecond())
	fmt.Println(d.MouthNextSecond())
	fmt.Println(d.YearNextSecond())
	fmt.Println(d.DayLastSecond())

	fmt.Println(d.DayOfWeek())
	fmt.Println(d.DayOfMouth())
	fmt.Println(d.DayOfYear())

	fmt.Println(d.SubYears(1))
	fmt.Println(d.SubMouths(1))
	fmt.Println(d.SubWeeks(1))
	fmt.Println(d.SubDays(1))
	fmt.Println(d.SubHours(1))
	fmt.Println(d.SubMinutes(10))
	fmt.Println(d.SubSeconds(160))

	fmt.Println(d.AddYears(10))
	fmt.Println(d.AddMouths(1))
	fmt.Println(d.AddWeeks(11))
	fmt.Println(d.AddDays(10))
	fmt.Println(d.AddHours(10))
	fmt.Println(d.AddMinutes(10))
	fmt.Println(d.AddSeconds(160))

}
