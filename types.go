package cronparse

import (
	"fmt"
	"strconv"
	"time"

	"github.com/winebarrel/cronparse/utils"
)

// number
type Number struct {
	Value int `@Number`
}

func (v *Number) String() string {
	return strconv.Itoa(v.Value)
}

func (v *Number) Match(x int) bool {
	return v.Value == x
}

// month name
type MonthName struct {
	Value string `@Month`
}

func (v *MonthName) String() string {
	return v.Value
}

func (v *MonthName) Match(x time.Month) bool {
	return utils.MonthNameToNumber(v.Value) == int(x)
}

// week name
type WeekName struct {
	Value string `@Week`
}

func (v *WeekName) String() string {
	return v.Value
}

func (v *WeekName) Match(x time.Weekday) bool {
	return utils.WeekNameToNumber(v.Value)%7 == int(x)
}

// number range
type NumberRange struct {
	From int `@Number`
	To   int `"-" @Number`
}

func (v *NumberRange) String() string {
	return fmt.Sprintf("%d-%d", v.From, v.To)
}

func (v *NumberRange) Match(x int) bool {
	return v.From <= x && x <= v.To
}

// week range
type WeekRange struct {
	From string `@Week`
	To   string `"-" @Week`
}

func (v *WeekRange) String() string {
	return fmt.Sprintf("%s-%s", v.From, v.To)
}

func (v *WeekRange) Match(x time.Weekday) bool {
	from := utils.WeekNameToNumber(v.From) % 7
	to := utils.WeekNameToNumber(v.To)
	n := int(x)

	if from != 0 && x == time.Sunday {
		n = 7
	}

	return from <= n && n <= to
}

// month range
type MonthRange struct {
	From string `@Month`
	To   string `"-" @Month`
}

func (v *MonthRange) String() string {
	return fmt.Sprintf("%s-%s", v.From, v.To)
}

func (v *MonthRange) Match(x time.Month) bool {
	from := utils.MonthNameToNumber(v.From)
	to := utils.MonthNameToNumber(v.To)
	return from <= int(x) && int(x) <= to
}

// all
type All struct {
	Value struct{} `"*"`
}

func (v *All) String() string {
	return "*"
}

func (v *All) Match(x interface{}) bool {
	return true
}

// increment
type Increment struct {
	Wildcard bool `( @"*"`
	Top      int  `| @Number )`
	Buttom   int  `"/" @Number`
}

func (v *Increment) String() string {
	if v.Wildcard {
		return fmt.Sprintf("*/%d", v.Buttom)
	} else {
		return fmt.Sprintf("%d/%d", v.Top, v.Buttom)
	}
}

func (v *Increment) Match(x int) bool {
	top := v.Top % v.Buttom

	if v.Wildcard {
		top = 0
	}

	return x >= v.Top && x%v.Buttom == top
}

// any
type Any struct {
	Value struct{} `"?"`
}

func (v *Any) String() string {
	return "?"
}

func (v *Any) Match(x interface{}) bool {
	return true
}

// last of month
type LastOfMonth struct {
	Value struct{} `"L"`
}

func (v *LastOfMonth) String() string {
	return "L"
}

func (v *LastOfMonth) Match(t time.Time) bool {
	return utils.LastOfMonth(t) == t.Day()
}

// last of week
type LastOfWeek struct {
	Value struct{} `"L"`
}

func (v *LastOfWeek) String() string {
	return "L"
}

func (v *LastOfWeek) Match(t time.Time) bool {
	return t.Weekday() == time.Saturday
}

// weekday
type Weekday struct {
	Value int `@Number "W"`
}

func (v *Weekday) String() string {
	return fmt.Sprintf("%dW", v.Value)
}

func (v *Weekday) Match(base time.Time) bool {
	t := time.Date(base.Year(), base.Month(), v.Value, 0, 0, 0, 0, time.UTC)
	return utils.NearestWeekday(t) == base.Day()
}

// instance
type Instance struct {
	DayOfWeek    int `@Number`
	NthDayOfWeek int `"#" @Number`
}

func (v *Instance) String() string {
	return fmt.Sprintf("%d#%d", v.DayOfWeek, v.NthDayOfWeek)
}

func (v *Instance) Match(t time.Time) bool {
	return utils.NthDayOfWeek(t, time.Weekday(v.DayOfWeek), v.NthDayOfWeek) == t.Day()
}
