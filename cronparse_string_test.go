package cronparse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronparse"
)

func TestNumberToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Number{Value: 1}
	assert.Equal(x.String(), "1")
}

func TestMonthNameToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.MonthName{Value: "JAN"}
	assert.Equal(x.String(), "JAN")
}

func TestWeekNameToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.WeekName{Value: "SUN"}
	assert.Equal(x.String(), "SUN")
}

func TestNumberRangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.NumberRange{From: 0, To: 59}
	assert.Equal(x.String(), "0-59")
}

func TestWeekRangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.WeekRange{From: "SUN", To: "SAT"}
	assert.Equal(x.String(), "SUN-SAT")
}

func TestMonthRangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.MonthRange{From: "JAN", To: "DEC"}
	assert.Equal(x.String(), "JAN-DEC")
}

func TestAllToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.All{}
	assert.Equal(x.String(), "*")
}

func TestIncrementToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Increment{Top: 0, Buttom: 5}
	assert.Equal(x.String(), "0/5")
}

func TestIncrementWildcardToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Increment{Wildcard: true, Buttom: 5}
	assert.Equal(x.String(), "*/5")
}

func TestAnyToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Any{}
	assert.Equal(x.String(), "?")
}

func TestLastToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Last{}
	assert.Equal(x.String(), "L")
}

func TestWeekdayToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Weekday{Value: 3}
	assert.Equal(x.String(), "3W")
}

func TestInstanceToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Instance{DayOfWeek: 6, NthDayOfWeek: 3}
	assert.Equal(x.String(), "6#3")
}

func TestMinutesExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.MinutesExp{
		CommonExp: cronparse.CommonExp{
			Number: &cronparse.Number{Value: 1},
		},
	}
	assert.Equal(x.String(), "1")
}

func TestMinutesToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Minutes{
		Exps: []*cronparse.MinutesExp{
			{
				cronparse.CommonExp{
					Number: &cronparse.Number{Value: 1},
				},
			},
			{
				cronparse.CommonExp{
					Number: &cronparse.Number{Value: 2},
				},
			},
		},
	}
	assert.Equal(x.String(), "1,2")
}

func TestHoursExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.HoursExp{
		CommonExp: cronparse.CommonExp{
			Number: &cronparse.Number{Value: 1},
		},
	}
	assert.Equal(x.String(), "1")
}

func TestHoursToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Hours{
		Exps: []*cronparse.HoursExp{
			{
				cronparse.CommonExp{
					Number: &cronparse.Number{Value: 1},
				},
			},
			{
				cronparse.CommonExp{
					Number: &cronparse.Number{Value: 2},
				},
			},
		},
	}
	assert.Equal(x.String(), "1,2")
}

func TestDayOfMonthExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.DayOfMonthExp{
		CommonExp: cronparse.CommonExp{
			Number: &cronparse.Number{Value: 1},
		},
	}
	assert.Equal(x.String(), "1")
}

func TestDayOfMonthToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.DayOfMonth{
		Exps: []*cronparse.DayOfMonthExp{
			{
				CommonExp: cronparse.CommonExp{
					Number: &cronparse.Number{Value: 1},
				},
			},
			{
				CommonExp: cronparse.CommonExp{
					Number: &cronparse.Number{Value: 2},
				},
			},
		},
	}
	assert.Equal(x.String(), "1,2")
}

func TestMonthExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.MonthExp{
		CommonExp: cronparse.CommonExp{
			Number: &cronparse.Number{Value: 1},
		},
	}
	assert.Equal(x.String(), "1")
}

func TestMonthToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Month{
		Exps: []*cronparse.MonthExp{
			{
				CommonExp: cronparse.CommonExp{
					Number: &cronparse.Number{Value: 1},
				},
			},
			{
				CommonExp: cronparse.CommonExp{
					Number: &cronparse.Number{Value: 2},
				},
			},
		},
	}
	assert.Equal(x.String(), "1,2")
}

func TestDayOfWeekExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.DayOfWeekExp{
		CommonExp: cronparse.CommonExp{
			Number: &cronparse.Number{Value: 1},
		},
	}
	assert.Equal(x.String(), "1")
}

func TestDayOfWeekToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.DayOfWeek{
		Exps: []*cronparse.DayOfWeekExp{
			{
				CommonExp: cronparse.CommonExp{
					Number: &cronparse.Number{Value: 1},
				},
			},
			{
				CommonExp: cronparse.CommonExp{
					Number: &cronparse.Number{Value: 2},
				},
			},
		},
	}
	assert.Equal(x.String(), "1,2")
}

func TestYearExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.YearExp{
		CommonExp: cronparse.CommonExp{
			Number: &cronparse.Number{Value: 1},
		},
	}
	assert.Equal(x.String(), "1")
}

func TestYearToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Year{
		Exps: []*cronparse.YearExp{
			{
				CommonExp: cronparse.CommonExp{
					Number: &cronparse.Number{Value: 1},
				},
			},
			{
				CommonExp: cronparse.CommonExp{
					Number: &cronparse.Number{Value: 2},
				},
			},
		},
	}
	assert.Equal(x.String(), "1,2")
}

func TestExpressionToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronparse.Expression{
		Minutes: &cronparse.Minutes{
			Exps: []*cronparse.MinutesExp{
				{
					CommonExp: cronparse.CommonExp{
						Number: &cronparse.Number{},
					},
				},
			},
		},
		Hours: &cronparse.Hours{
			Exps: []*cronparse.HoursExp{
				{
					CommonExp: cronparse.CommonExp{
						Number: &cronparse.Number{
							Value: 10,
						},
					},
				},
			},
		},
		DayOfMonth: &cronparse.DayOfMonth{
			Exps: []*cronparse.DayOfMonthExp{
				{
					CommonExp: cronparse.CommonExp{
						All: &cronparse.All{
							Value: struct{}{},
						},
					},
				},
			},
		},
		Month: &cronparse.Month{
			Exps: []*cronparse.MonthExp{
				{
					CommonExp: cronparse.CommonExp{
						All: &cronparse.All{
							Value: struct{}{},
						},
					},
				},
			},
		},
		DayOfWeek: &cronparse.DayOfWeek{
			Exps: []*cronparse.DayOfWeekExp{
				{
					CommonExp: cronparse.CommonExp{},
					Any: &cronparse.Any{
						Value: struct{}{},
					},
				},
			},
		},
		Year: &cronparse.Year{
			Exps: []*cronparse.YearExp{
				{
					CommonExp: cronparse.CommonExp{
						All: &cronparse.All{
							Value: struct{}{},
						},
					},
				},
			},
		},
	}
	assert.Equal(x.String(), "0 10 * * ? *")
}
