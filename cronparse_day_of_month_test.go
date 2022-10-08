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
	assert.Equal(&cronparse.All{}, cron.DayOfMonth.Exps[0].All)
}

func TestDayOfMonthNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * 1 * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Number{Value: 1}, cron.DayOfMonth.Exps[0].Number)
}

func TestDayOfMonthNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * 1-30 * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.NumberRange{From: 1, To: 30}, cron.DayOfMonth.Exps[0].NumberRange)
}

func TestDayOfMonthIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * 1/5 * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Top: 1, Buttom: 5}, cron.DayOfMonth.Exps[0].Increment)
}

func TestDayOfMonthIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * */5 * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 5}, cron.DayOfMonth.Exps[0].Increment)
}

func TestDayOfMonthAny(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * * *")
	assert.NoError(err)
	assert.Equal(&cronparse.Any{}, cron.DayOfMonth.Exps[0].Any)
}

func TestDayOfMonthLast(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * L * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Last{}, cron.DayOfMonth.Exps[0].Last)
}

func TestDayOfMonthWeekday(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * 3W * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Weekday{Value: 3}, cron.DayOfMonth.Exps[0].Weekday)
}

func TestDayOfMonthComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * *,1,1-30,1/5,*/5,?,L,3W * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.All{}, cron.DayOfMonth.Exps[0].All)
	assert.Equal(&cronparse.Number{Value: 1}, cron.DayOfMonth.Exps[1].Number)
	assert.Equal(&cronparse.NumberRange{From: 1, To: 30}, cron.DayOfMonth.Exps[2].NumberRange)
	assert.Equal(&cronparse.Increment{Top: 1, Buttom: 5}, cron.DayOfMonth.Exps[3].Increment)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 5}, cron.DayOfMonth.Exps[4].Increment)
	assert.Equal(&cronparse.Any{}, cron.DayOfMonth.Exps[5].Any)
	assert.Equal(&cronparse.Last{}, cron.DayOfMonth.Exps[6].Last)
	assert.Equal(&cronparse.Weekday{Value: 3}, cron.DayOfMonth.Exps[7].Weekday)
}
