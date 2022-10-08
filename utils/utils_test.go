package utils_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronparse/utils"
)

func TestWeekNameToNumber(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		mon      string
		expected int
	}{
		{"mon", 1},
		{"tue", 2},
		{"wed", 3},
		{"thu", 4},
		{"fri", 5},
		{"sat", 6},
		{"sun", 7},
		{"xxx", -1},
	}

	for _, t := range tt {
		assert.Equal(t.expected, utils.WeekNameToNumber(t.mon), t.mon)
	}
}

func TestMonthNameToNumber(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		mon      string
		expected int
	}{
		{"jan", 1},
		{"feb", 2},
		{"mar", 3},
		{"apr", 4},
		{"may", 5},
		{"jun", 6},
		{"jul", 7},
		{"aug", 8},
		{"sep", 9},
		{"oct", 10},
		{"nov", 11},
		{"dec", 12},
		{"xxx", -1},
	}

	for _, t := range tt {
		assert.Equal(t.expected, utils.MonthNameToNumber(t.mon), t.mon)
	}
}

func TestLastOfMonth(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		expected int
	}{
		{time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 2, 1, 9, 0, 0, 0, time.UTC), 28},
		{time.Date(2023, 3, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 4, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 5, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 6, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 7, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 8, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 9, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 10, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 11, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 12, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 2, 1, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2024, 3, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 4, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 5, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 6, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 7, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 8, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 9, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 10, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 11, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 12, 1, 9, 0, 0, 0, time.UTC), 31},
	}

	for _, t := range tt {
		assert.Equal(t.expected, utils.LastOfMonth(t.tm), t.tm)
	}
}

func TestNearestWeekday(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		expected int
	}{
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), 3},
		{time.Date(2022, 11, 4, 9, 0, 0, 0, time.UTC), 4},
		{time.Date(2022, 11, 5, 9, 0, 0, 0, time.UTC), 4},
		{time.Date(2022, 11, 6, 9, 0, 0, 0, time.UTC), 7},
		{time.Date(2022, 11, 7, 9, 0, 0, 0, time.UTC), 7},
		{time.Date(2022, 11, 8, 9, 0, 0, 0, time.UTC), 8},
	}

	for _, t := range tt {
		assert.Equal(t.expected, utils.NearestWeekday(t.tm), t.tm)
	}
}

func TestNthDayOfWeek(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		w        time.Weekday
		nth      int
		expected int
	}{
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Monday, 1, 3},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Monday, 2, 10},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Monday, 3, 17},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Monday, 4, 24},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Monday, 5, 31},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 1, 4},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 2, 11},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 3, 18},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 4, 25},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 1, 5},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 2, 12},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 3, 19},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 4, 26},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 1, 6},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 2, 13},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 3, 20},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 4, 27},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Friday, 1, 7},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Friday, 2, 14},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Friday, 3, 21},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Friday, 4, 28},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 1, 1},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 2, 8},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 3, 15},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 4, 22},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 5, 29},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 1, 2},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 2, 9},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 3, 16},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 4, 23},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 5, 30},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Monday, 1, 7},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Monday, 2, 14},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Monday, 3, 21},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Monday, 4, 28},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 1, 1},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 2, 8},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 3, 15},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 4, 22},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 5, 29},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 1, 2},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 2, 9},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 3, 16},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 4, 23},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 5, 30},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 1, 3},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 2, 10},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 3, 17},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 4, 24},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Friday, 1, 4},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Friday, 2, 11},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Friday, 3, 18},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Friday, 4, 25},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 1, 5},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 2, 12},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 3, 19},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 4, 26},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 1, 6},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 2, 13},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 3, 20},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 4, 27},
	}

	for _, t := range tt {
		assert.Equal(t.expected, utils.NthDayOfWeek(t.tm, t.w, t.nth), t)
	}
}
