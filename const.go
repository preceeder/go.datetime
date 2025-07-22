package datetime

import "time"

// max constants
const (
	MaxYear       = 9999
	MaxMonth      = 12
	MaxDay        = 31
	MaxHour       = 23
	MaxMinute     = 59
	MaxSecond     = 59
	MaxNanosecond = 999999999
)

// min constants
const (
	MinYear       = 1
	MinMonth      = 1
	MinDay        = 1
	MinHour       = 0
	MinMinute     = 0
	MinSecond     = 0
	MinNanosecond = 0
)

const (
	DaysPerWeek = 7
)

var (
	// DefaultWeekStartsAt default start date of the week
	DefaultWeekStartsAt = time.Monday

	// DefaultWeekendDays default weekend days of the week
	DefaultWeekendDays = []time.Weekday{
		time.Saturday, time.Sunday,
	}
)
