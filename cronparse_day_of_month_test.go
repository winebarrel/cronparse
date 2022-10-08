package cronparse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronparse"
)

func TestDayOfMonthAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? *")
	assert.NoError(err)
	assert.Equal(cron.DayOfMonth.Exps[0].All, &cronparse.All{})
}

func TestDayOfMonthNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * 1 * ? *")
	assert.NoError(err)
	assert.Equal(cron.DayOfMonth.Exps[0].Number, &cronparse.Number{Value: 1})
}

func TestDayOfMonthNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * 1-30 * ? *")
	assert.NoError(err)
	assert.Equal(cron.DayOfMonth.Exps[0].NumberRange, &cronparse.NumberRange{From: 1, To: 30})
}

func TestDayOfMonthIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * 1/5 * ? *")
	assert.NoError(err)
	assert.Equal(cron.DayOfMonth.Exps[0].Increment, &cronparse.Increment{Top: 1, Buttom: 5})
}

func TestDayOfMonthIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * */5 * ? *")
	assert.NoError(err)
	assert.Equal(cron.DayOfMonth.Exps[0].Increment, &cronparse.Increment{Wildcard: true, Buttom: 5})
}

func TestDayOfMonthAny(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * * *")
	assert.NoError(err)
	assert.Equal(cron.DayOfMonth.Exps[0].Any, &cronparse.Any{})
}

func TestDayOfMonthLast(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * L * ? *")
	assert.NoError(err)
	assert.Equal(cron.DayOfMonth.Exps[0].Last, &cronparse.Last{})
}

func TestDayOfMonthWeekday(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * 3W * ? *")
	assert.NoError(err)
	assert.Equal(cron.DayOfMonth.Exps[0].Weekday, &cronparse.Weekday{Value: 3})
}

func TestDayOfMonthComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * *,1,1-30,1/5,*/5,?,L,3W * ? *")
	assert.NoError(err)
	assert.Equal(cron.DayOfMonth.Exps[0].All, &cronparse.All{})
	assert.Equal(cron.DayOfMonth.Exps[1].Number, &cronparse.Number{Value: 1})
	assert.Equal(cron.DayOfMonth.Exps[2].NumberRange, &cronparse.NumberRange{From: 1, To: 30})
	assert.Equal(cron.DayOfMonth.Exps[3].Increment, &cronparse.Increment{Top: 1, Buttom: 5})
	assert.Equal(cron.DayOfMonth.Exps[4].Increment, &cronparse.Increment{Wildcard: true, Buttom: 5})
	assert.Equal(cron.DayOfMonth.Exps[5].Any, &cronparse.Any{})
	assert.Equal(cron.DayOfMonth.Exps[6].Last, &cronparse.Last{})
	assert.Equal(cron.DayOfMonth.Exps[7].Weekday, &cronparse.Weekday{Value: 3})
}
