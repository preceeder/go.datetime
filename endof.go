package datetime

import "time"

func (t *DateTime) StartOfMinute(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	return time.Date(t.time.Year(), t.time.Month(), t.time.Day(), t.time.Hour(), t.time.Minute(), MinSecond, MinNanosecond, timeZoneHandler(zone...))

}

func (t *DateTime) EndOfMinute(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	return time.Date(t.time.Year(), t.time.Month(), t.time.Day(), t.time.Hour(), t.time.Minute(), MaxSecond, MaxNanosecond, timeZoneHandler(zone...))
}

func (t *DateTime) StartOfHour(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	return time.Date(t.time.Year(), t.time.Month(), t.time.Day(), t.time.Hour(), MinMinute, MinSecond, MinNanosecond, timeZoneHandler(zone...))
}

func (t *DateTime) EndOfHour(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	return time.Date(t.time.Year(), t.time.Month(), t.time.Day(), t.time.Hour(), MaxMinute, MaxSecond, MaxNanosecond, timeZoneHandler(zone...))
}

func (t *DateTime) StartOfDay(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	return time.Date(t.time.Year(), t.time.Month(), t.time.Day(), MinHour, MinMinute, MinSecond, MinNanosecond, timeZoneHandler(zone...))
}

func (t *DateTime) EndOfDay(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	return time.Date(t.time.Year(), t.time.Month(), t.time.Day(), MaxHour, MaxMinute, MaxSecond, MaxNanosecond, timeZoneHandler(zone...))
}

func (t *DateTime) StartOfWeek(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	dayOfWeek, weekStartsAt := t.time.Weekday(), t.weekStartsAt
	if dayOfWeek == weekStartsAt {
		return t.StartOfDay()
	}
	return t.Copy().SubDays(int(DaysPerWeek+dayOfWeek-weekStartsAt) % DaysPerWeek).StartOfDay()
}

func (t *DateTime) EndOfWeek(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	dayOfWeek, weekEndsAt := t.time.Weekday(), t.WeekEndsAt()
	if dayOfWeek == weekEndsAt {
		return t.EndOfDay()
	}
	return t.Copy().AddDays(int(DaysPerWeek-dayOfWeek+weekEndsAt) % DaysPerWeek).EndOfDay()
}

func (t *DateTime) StartOfMonth(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	return time.Date(t.time.Year(), t.time.Month(), MinDay, MinHour, MinMinute, MinSecond, MinNanosecond, timeZoneHandler(zone...))
}

func (t *DateTime) EndOfMonth(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	return time.Date(t.time.Year(), t.time.Month()+1, 0, MaxHour, MaxMinute, MaxSecond, MaxNanosecond, timeZoneHandler(zone...))
}

func (t *DateTime) StartOfYear(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	return time.Date(t.time.Year(), MinMonth, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond, timeZoneHandler(zone...))
}

func (t *DateTime) EndOfYear(zone ...string) time.Time {
	if t.time.IsZero() {
		t = Now(zone...)
	}
	return time.Date(t.time.Year(), MaxMonth, MaxDay, MaxHour, MaxMinute, MaxSecond, MaxNanosecond, timeZoneHandler(zone...))
}
