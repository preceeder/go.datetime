package datetime

import "math"

func (t *DateTime) DiffYears(date ...*DateTime) int64 {
	start := t.time
	end := t.Now().time
	if len(date) > 0 {
		end = date[0].time
	}
	dy, dm, dd := end.Year()-start.Year(), end.Month()-start.Month(), end.Day()-start.Day()
	if dm < 0 || (dm == 0 && dd < 0) {
		dy--
	}
	if dy < 0 && (dd != 0 || dm != 0) {
		dy++
	}
	return int64(dy)
}

func (t *DateTime) DiffMouths(date ...*DateTime) int64 {
	start := t.time
	end := t.Now().time
	if len(date) > 0 {
		end = date[0].time
	}
	startYear, startMonth, _ := start.Date()
	endYear, endMonth, _ := end.Date()
	diffYear, diffMonth := endYear-startYear, endMonth-startMonth
	return int64(diffYear*MonthsPerYear + int(diffMonth))
}

func (t *DateTime) DiffWeeks(date ...*DateTime) int64 {
	start := t.time
	end := t.Now().time
	if len(date) > 0 {
		end = date[0].time
	}
	return int64(math.Floor(float64((end.Unix() - start.Unix()) / (7 * 24 * 3600))))
}

func (t *DateTime) DiffDays(date ...*DateTime) int64 {
	start := t.time
	end := t.Now().time
	if len(date) > 0 {
		end = date[0].time
	}
	return int64(math.Floor(float64((end.Unix() - start.Unix()) / (24 * 3600))))
}

func (t *DateTime) DiffHours(date ...*DateTime) int64 {
	return t.DiffSeconds(date...) / 3600
}

func (t *DateTime) DiffMinutes(date ...*DateTime) int64 {
	return t.DiffSeconds(date...) / 60
}

func (t *DateTime) DiffSeconds(date ...*DateTime) int64 {
	start := t.time
	end := t.Now().time
	if len(date) > 0 {
		end = date[0].time
	}
	return end.Unix() - start.Unix()
}
