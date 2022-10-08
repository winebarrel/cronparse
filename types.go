package cronparse

import (
	"fmt"
	"strconv"
)

// number
type Number struct {
	Value int `@Number`
}

func (v *Number) String() string {
	return strconv.Itoa(v.Value)
}

// month name
type MonthName struct {
	Value string `@Month`
}

func (v *MonthName) String() string {
	return v.Value
}

// week name
type WeekName struct {
	Value string `@Week`
}

func (v *WeekName) String() string {
	return v.Value
}

// number range
type NumberRange struct {
	From int `@Number`
	To   int `"-" @Number`
}

func (v *NumberRange) String() string {
	return fmt.Sprintf("%d-%d", v.From, v.To)
}

// week range
type WeekRange struct {
	From string `@Week`
	To   string `"-" @Week`
}

func (v *WeekRange) String() string {
	return fmt.Sprintf("%s-%s", v.From, v.To)
}

// month range
type MonthRange struct {
	From string `@Month`
	To   string `"-" @Month`
}

func (v *MonthRange) String() string {
	return fmt.Sprintf("%s-%s", v.From, v.To)
}

// all
type All struct {
	Value struct{} `"*"`
}

func (v *All) String() string {
	return "*"
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

// any
type Any struct {
	Value struct{} `"?"`
}

func (v *Any) String() string {
	return "?"
}

// last
type Last struct {
	Value struct{} `"L"`
}

func (v *Last) String() string {
	return "L"
}

// weekday
type Weekday struct {
	Value int `@Number "W"`
}

func (v *Weekday) String() string {
	return fmt.Sprintf("%dW", v.Value)
}

// instance
type Instance struct {
	DayOfWeek    int `@Number`
	NthDayOfWeek int `"#" @Number`
}

func (v *Instance) String() string {
	return fmt.Sprintf("%d#%d", v.DayOfWeek, v.NthDayOfWeek)
}
