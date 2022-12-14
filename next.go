package cronparse

import (
	"time"
)

func (v *Expression) Next(from time.Time) time.Time {
	schedule := v.NextN(from, 1)

	if len(schedule) == 0 {
		return time.Time{}
	}

	return schedule[0]
}

func (v *Expression) NextN(from time.Time, n int) []time.Time {
	schedule := []time.Time{}
	years := v.candidateYears(from)

	if len(years) == 0 {
		return schedule
	}

	months := v.candidateMonths(from)

	if len(months) == 0 {
		return schedule
	}

	hours := v.candidateHours(from)

	if len(hours) == 0 {
		return schedule
	}

	minutes := v.candidateMinutes(from)

	if len(minutes) == 0 {
		return schedule
	}

	var DayMatch func(time.Time) bool

	if !v.DayOfMonth.HasAny() && v.DayOfWeek.HasAny() {
		DayMatch = v.DayOfMonth.Match
	} else if v.DayOfMonth.HasAny() && !v.DayOfWeek.HasAny() {
		DayMatch = v.DayOfWeek.Match
	} else {
		return schedule
	}

YEAR:
	for _, year := range years {
		for _, month := range months {
			if year == from.Year() && month < from.Month() {
				continue
			}

			for day := 1; day <= 31; day++ {
				if year == from.Year() && month == from.Month() && day < from.Day() {
					continue
				}

				dayOfMonth := time.Date(year, time.Month(month), day, 0, 0, 0, 0, from.Location())

				if dayOfMonth.Month() != month {
					break
				}

				if !DayMatch(dayOfMonth) {
					continue
				}

				for _, hour := range hours {
					if year == from.Year() && month == from.Month() && day == from.Day() && hour < from.Hour() {
						continue
					}

					for _, minute := range minutes {
						if year == from.Year() && month == from.Month() && day == from.Day() && hour == from.Hour() && minute < from.Minute() {
							continue
						}

						schedule = append(schedule, time.Date(year, time.Month(month), day, hour, minute, 0, 0, from.Location()))

						if len(schedule) >= n {
							break YEAR
						}
					}
				}
			}
		}
	}

	return schedule
}

func (v *Expression) candidateYears(from time.Time) []int {
	candidates := []int{}

	for year := from.Year(); year <= 2199; year++ {
		t := time.Date(year, 1, 1, 0, 0, 0, 0, from.Location())

		if v.Year.Match(t) {
			candidates = append(candidates, year)
		}
	}

	return candidates
}

func (v *Expression) candidateMonths(from time.Time) []time.Month {
	candidates := []time.Month{}

	for month := time.January; month <= time.December; month++ {
		t := time.Date(from.Year(), month, 1, 0, 0, 0, 0, from.Location())

		if v.Month.Match(t) {
			candidates = append(candidates, month)
		}
	}

	return candidates
}

func (v *Expression) candidateHours(from time.Time) []int {
	candidates := []int{}

	for hour := 0; hour <= 23; hour++ {
		t := time.Date(from.Year(), from.Month(), from.Day(), hour, 0, 0, 0, from.Location())

		if v.Hours.Match(t) {
			candidates = append(candidates, hour)
		}
	}

	return candidates
}

func (v *Expression) candidateMinutes(from time.Time) []int {
	candidates := []int{}

	for minute := 0; minute <= 59; minute++ {
		t := time.Date(from.Year(), from.Month(), from.Day(), from.Hour(), minute, 0, 0, from.Location())

		if v.Minutes.Match(t) {
			candidates = append(candidates, minute)
		}
	}

	return candidates
}
