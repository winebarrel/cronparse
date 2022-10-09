package cronparse_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronparse"
)

func TestMatchNumber(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Number{Value: 1}
	assert.True(x.Match(1))
}

func TestMatchMonthName(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.MonthName{Value: "JAN"}
	assert.True(x.Match(time.January))
}

func TestMatchWeekName(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.WeekName{Value: "MON"}
	assert.True(x.Match(time.Monday))
}

func TestMatchWeekNameSunday(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.WeekName{Value: "SUN"}
	assert.True(x.Match(time.Sunday))
}

func TestMatchNumberRange(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.NumberRange{From: 1, To: 3}

	tt := []struct {
		num      int
		expected bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, true},
		{4, false},
	}

	for _, t := range tt {
		assert.Equal(t.expected, x.Match(t.num))
	}
}

func TestMatchWeekRange1(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.WeekRange{From: "SUN", To: "FRI"}

	tt := []struct {
		w        time.Weekday
		expected bool
	}{
		{time.Sunday, true},
		{time.Monday, true},
		{time.Tuesday, true},
		{time.Wednesday, true},
		{time.Thursday, true},
		{time.Friday, true},
		{time.Saturday, false},
	}

	for _, t := range tt {
		assert.Equal(t.expected, x.Match(t.w), t.w)
	}
}

func TestMatchWeekRange2(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.WeekRange{From: "TUE", To: "SUN"}

	tt := []struct {
		w        time.Weekday
		expected bool
	}{
		{time.Sunday, true},
		{time.Monday, false},
		{time.Tuesday, true},
		{time.Wednesday, true},
		{time.Thursday, true},
		{time.Friday, true},
		{time.Saturday, true},
	}

	for _, t := range tt {
		assert.Equal(t.expected, x.Match(t.w), t.w)
	}
}

func TestMatchMonthRange(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.MonthRange{From: "FEB", To: "NOV"}

	tt := []struct {
		name     time.Month
		expected bool
	}{

		{time.January, false},
		{time.February, true},
		{time.March, true},
		{time.April, true},
		{time.May, true},
		{time.June, true},
		{time.July, true},
		{time.August, true},
		{time.September, true},
		{time.October, true},
		{time.November, true},
		{time.December, false},
	}

	for _, t := range tt {
		assert.Equal(t.expected, x.Match(t.name), t.name)
	}
}

func TestMatchAll(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.All{}
	assert.True(x.Match("xxx"))
}

func TestMatchIncrement1(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Increment{Top: 1, Buttom: 3}

	tt := []struct {
		num      int
		expected bool
	}{
		{0, false},
		{1, true},
		{2, false},
		{3, false},
		{4, true},
		{5, false},
		{6, false},
		{7, true},
		{8, false},
	}

	for _, t := range tt {
		assert.Equal(t.expected, x.Match(t.num, 0), t.num)
	}
}

func TestMatchIncrement2(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Increment{Top: 5, Buttom: 3}

	tt := []struct {
		num      int
		expected bool
	}{
		{0, false},
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{5, true},
		{6, false},
		{7, false},
		{8, true},
	}

	for _, t := range tt {
		assert.Equal(t.expected, x.Match(t.num, 0), t.num)
	}
}

func TestMatchIncrementWildcard1(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Increment{Wildcard: true, Buttom: 3}

	tt := []struct {
		num      int
		expected bool
	}{
		{0, true},
		{1, false},
		{2, false},
		{3, true},
		{4, false},
		{5, false},
		{6, true},
		{7, false},
		{8, false},
	}

	for _, t := range tt {
		assert.Equal(t.expected, x.Match(t.num, 0), t.num)
	}
}

func TestMatchIncrementWildcard2(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Increment{Wildcard: true, Buttom: 3}

	tt := []struct {
		num      int
		expected bool
	}{
		{0, false},
		{1, true},
		{2, false},
		{3, false},
		{4, true},
		{5, false},
		{6, false},
		{7, true},
		{8, false},
	}

	for _, t := range tt {
		assert.Equal(t.expected, x.Match(t.num, 1), t.num)
	}
}

func TestMatchAny(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Any{}
	assert.True(x.Match("xxx"))
}

func TestMatchWeekday(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		w        *cronparse.Weekday
		expected bool
	}{
		{time.Date(2022, 11, 4, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 4}, true},
		{time.Date(2022, 11, 4, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 5}, true},
		{time.Date(2022, 11, 4, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 6}, false},
		{time.Date(2022, 11, 4, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 7}, false},
		{time.Date(2022, 11, 5, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 4}, false},
		{time.Date(2022, 11, 5, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 5}, false},
		{time.Date(2022, 11, 5, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 6}, false},
		{time.Date(2022, 11, 5, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 7}, false},
		{time.Date(2022, 11, 6, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 4}, false},
		{time.Date(2022, 11, 6, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 5}, false},
		{time.Date(2022, 11, 6, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 6}, false},
		{time.Date(2022, 11, 6, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 7}, false},
		{time.Date(2022, 11, 7, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 4}, false},
		{time.Date(2022, 11, 7, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 5}, false},
		{time.Date(2022, 11, 7, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 6}, true},
		{time.Date(2022, 11, 7, 9, 0, 0, 0, time.UTC), &cronparse.Weekday{Value: 7}, true},
	}

	for _, t := range tt {
		assert.Equal(t.expected, t.w.Match(t.tm), t)
	}
}

func TestMatchInstance(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		i        *cronparse.Instance
		expected bool
	}{
		{time.Date(2022, 11, 6, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Monday), NthDayOfWeek: 1}, false},
		{time.Date(2022, 11, 7, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Monday), NthDayOfWeek: 1}, true},
		{time.Date(2022, 11, 8, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Monday), NthDayOfWeek: 1}, false},
		{time.Date(2022, 11, 13, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Monday), NthDayOfWeek: 2}, false},
		{time.Date(2022, 11, 14, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Monday), NthDayOfWeek: 2}, true},
		{time.Date(2022, 11, 15, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Monday), NthDayOfWeek: 2}, false},
		{time.Date(2022, 11, 20, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Monday), NthDayOfWeek: 3}, false},
		{time.Date(2022, 11, 21, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Monday), NthDayOfWeek: 3}, true},
		{time.Date(2022, 11, 22, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Monday), NthDayOfWeek: 3}, false},
		{time.Date(2022, 11, 7, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Tuesday), NthDayOfWeek: 1}, false},
		{time.Date(2022, 11, 1, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Tuesday), NthDayOfWeek: 1}, true},
		{time.Date(2022, 11, 2, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Tuesday), NthDayOfWeek: 1}, false},
		{time.Date(2022, 11, 7, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Tuesday), NthDayOfWeek: 2}, false},
		{time.Date(2022, 11, 8, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Tuesday), NthDayOfWeek: 2}, true},
		{time.Date(2022, 11, 9, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Tuesday), NthDayOfWeek: 2}, false},
		{time.Date(2022, 11, 14, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Tuesday), NthDayOfWeek: 3}, false},
		{time.Date(2022, 11, 15, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Tuesday), NthDayOfWeek: 3}, true},
		{time.Date(2022, 11, 16, 9, 0, 0, 0, time.UTC), &cronparse.Instance{DayOfWeek: int(time.Tuesday), NthDayOfWeek: 3}, false},
	}

	for _, t := range tt {
		assert.Equal(t.expected, t.i.Match(t.tm), t)
	}
}

func TestMatchLastOfMonth(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		l        *cronparse.LastOfMonth
		expected bool
	}{
		{time.Date(2023, 1, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 2, 28, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 3, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 4, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 5, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 6, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 7, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 8, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 9, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 10, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 11, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 12, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 1, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 2, 29, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 3, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 4, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 5, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 6, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 7, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 8, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 9, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 10, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 11, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2024, 12, 31, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, true},
		{time.Date(2023, 1, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2023, 2, 27, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2023, 3, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2023, 4, 29, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2023, 5, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2023, 6, 29, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2023, 7, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2023, 8, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2023, 9, 29, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2023, 10, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2023, 11, 29, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2023, 12, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 1, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 2, 28, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 3, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 4, 29, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 5, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 6, 29, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 7, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 8, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 9, 29, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 10, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 11, 29, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
		{time.Date(2024, 12, 30, 9, 0, 0, 0, time.UTC), &cronparse.LastOfMonth{}, false},
	}

	for _, t := range tt {
		assert.Equal(t.expected, t.l.Match(t.tm), t.tm)
	}
}

func TestMatchLastOfWeek(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		l        *cronparse.LastOfWeek
		expected bool
	}{
		{time.Date(2023, 9, 1, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
		{time.Date(2023, 9, 2, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, true},
		{time.Date(2023, 9, 3, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
		{time.Date(2023, 9, 4, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
		{time.Date(2023, 9, 5, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
		{time.Date(2023, 9, 6, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
		{time.Date(2023, 9, 7, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
		{time.Date(2023, 9, 8, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
		{time.Date(2023, 9, 9, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, true},
		{time.Date(2023, 9, 10, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
		{time.Date(2023, 9, 11, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
		{time.Date(2023, 9, 12, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
		{time.Date(2023, 9, 13, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
		{time.Date(2023, 9, 14, 9, 0, 0, 0, time.UTC), &cronparse.LastOfWeek{}, false},
	}

	for _, t := range tt {
		assert.Equal(t.expected, t.l.Match(t.tm), t.tm)
	}
}
