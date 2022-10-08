package utils

import (
	"strings"
	"time"
)

var (
	MonthNames = []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
	WeekNames  = []string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}
)

func WeekNameToNumber(s string) int {
	for i, n := range WeekNames {
		if strings.EqualFold(n, s) {
			return i + 1
		}
	}

	return -1
}

func MonthNameToNumber(s string) int {
	for i, n := range MonthNames {
		if strings.EqualFold(n, s) {
			return i + 1
		}
	}

	return -1
}

func LastOfMonth(t time.Time) int {
	return t.AddDate(0, 1, -t.Day()).Day()
}

func NearestWeekday(t time.Time) int {
	if t.Weekday() == time.Saturday {
		return t.AddDate(0, 0, -1).Day()
	} else if t.Weekday() == time.Sunday {
		return t.AddDate(0, 0, 1).Day()
	}

	return t.Day()
}

func NthDayOfWeek(t time.Time, w time.Weekday, nth int) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	offset := (w + 7 - firstOfMonth.Weekday()) % 7
	nthDoW := firstOfMonth.AddDate(0, 0, 7*(nth-1)+int(offset))
	return nthDoW.Day()
}
