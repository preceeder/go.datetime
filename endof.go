package datetime

import "time"

func (t *DateTime) StartOfMinute() *DateTime {

	return &DateTime{
		weekStartsAt: t.weekStartsAt,
		location:     t.location,
		time:         time.Date(t.time.Year(), t.time.Month(), t.time.Day(), t.time.Hour(), t.time.Minute(), MinSecond, MinNanosecond, t.location),
	}

}

func (t *DateTime) EndOfMinute() *DateTime {
	return &DateTime{
		weekStartsAt: t.weekStartsAt,
		location:     t.location,
		time:         time.Date(t.time.Year(), t.time.Month(), t.time.Day(), t.time.Hour(), t.time.Minute(), MaxSecond, MaxNanosecond, t.location),
	}
}

func (t *DateTime) StartOfHour() *DateTime {
	return &DateTime{
		weekStartsAt: t.weekStartsAt,
		location:     t.location,
		time:         time.Date(t.time.Year(), t.time.Month(), t.time.Day(), t.time.Hour(), MinMinute, MinSecond, MinNanosecond, t.location),
	}
}

func (t *DateTime) EndOfHour() *DateTime {
	return &DateTime{
		weekStartsAt: t.weekStartsAt,
		location:     t.location,
		time:         time.Date(t.time.Year(), t.time.Month(), t.time.Day(), t.time.Hour(), MaxMinute, MaxSecond, MaxNanosecond, t.location),
	}
}

func (t *DateTime) StartOfDay(zone ...string) *DateTime {

	return &DateTime{
		weekStartsAt: t.weekStartsAt,
		location:     t.location,
		time:         time.Date(t.time.Year(), t.time.Month(), t.time.Day(), MinHour, MinMinute, MinSecond, MinNanosecond, t.location),
	}
}

func (t *DateTime) EndOfDay() *DateTime {
	return &DateTime{
		weekStartsAt: t.weekStartsAt,
		location:     t.location,
		time:         time.Date(t.time.Year(), t.time.Month(), t.time.Day(), MaxHour, MaxMinute, MaxSecond, MaxNanosecond, t.location),
	}
}

func (t *DateTime) StartOfWeek() *DateTime {
	dayOfWeek, weekStartsAt := t.time.Weekday(), t.weekStartsAt
	if dayOfWeek == weekStartsAt {
		return t.StartOfDay()
	}
	return t.Copy().SubDays(int(DaysPerWeek+dayOfWeek-weekStartsAt) % DaysPerWeek).StartOfDay()
}

func (t *DateTime) EndOfWeek() *DateTime {

	dayOfWeek, weekEndsAt := t.time.Weekday(), t.WeekEndsAt()
	if dayOfWeek == weekEndsAt {
		return t.EndOfDay()
	}
	return t.Copy().AddDays(int(DaysPerWeek-dayOfWeek+weekEndsAt) % DaysPerWeek).EndOfDay()
}

func (t *DateTime) StartOfMonth() *DateTime {
	return &DateTime{
		weekStartsAt: t.weekStartsAt,
		location:     t.location,
		time:         time.Date(t.time.Year(), t.time.Month(), MinDay, MinHour, MinMinute, MinSecond, MinNanosecond, t.location),
	}
}

func (t *DateTime) EndOfMonth() *DateTime {
	return &DateTime{
		weekStartsAt: t.weekStartsAt,
		location:     t.location,
		time:         time.Date(t.time.Year(), t.time.Month()+1, 0, MaxHour, MaxMinute, MaxSecond, MaxNanosecond, t.location),
	}
}

func (t *DateTime) StartOfYear() *DateTime {
	return &DateTime{
		weekStartsAt: t.weekStartsAt,
		location:     t.location,
		time:         time.Date(t.time.Year(), MinMonth, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond, t.location),
	}
}

func (t *DateTime) EndOfYear() *DateTime {
	return &DateTime{
		weekStartsAt: t.weekStartsAt,
		location:     t.location,
		time:         time.Date(t.time.Year(), MaxMonth, MaxDay, MaxHour, MaxMinute, MaxSecond, MaxNanosecond, t.location),
	}
}
