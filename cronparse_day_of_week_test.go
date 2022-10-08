package cronparse_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronparse"
)

func TestDayOfWeekAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * * *")
	assert.NoError(err)
	assert.Equal(cron.DayOfWeek.Exps[0].All, &cronparse.All{})
}

func TestDayOfWeekNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * 1 *")
	assert.NoError(err)
	assert.Equal(cron.DayOfWeek.Exps[0].Number, &cronparse.Number{Value: 1})
}

func TestDayOfWeekNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * 1-7 *")
	assert.NoError(err)
	assert.Equal(cron.DayOfWeek.Exps[0].NumberRange, &cronparse.NumberRange{From: 1, To: 7})
}

func TestDayOfWeekIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * 1/5 *")
	assert.NoError(err)
	assert.Equal(cron.DayOfWeek.Exps[0].Increment, &cronparse.Increment{Top: 1, Buttom: 5})
}

func TestDayOfWeekIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * */5 *")
	assert.NoError(err)
	assert.Equal(cron.DayOfWeek.Exps[0].Increment, &cronparse.Increment{Wildcard: true, Buttom: 5})
}

func TestDayOfWeekAny(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? *")
	assert.NoError(err)
	assert.Equal(cron.DayOfWeek.Exps[0].Any, &cronparse.Any{})
}

func TestDayOfWeekLast(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * L *")
	assert.NoError(err)
	assert.Equal(cron.DayOfWeek.Exps[0].Last, &cronparse.Last{})
}

func TestDayOfWeekString(t *testing.T) {
	assert := assert.New(t)
	tt := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}

	for _, t := range tt {
		cron, err := cronparse.Parser.ParseString("", fmt.Sprintf("* * ? * %s *", t))
		assert.NoError(err)
		assert.Equal(cron.DayOfWeek.Exps[0].String, &cronparse.WeekValue{Value: t})

		cron, err = cronparse.Parser.ParseString("", fmt.Sprintf("* * ? * %s *", strings.ToLower(t)))
		assert.NoError(err)
		assert.Equal(cron.DayOfWeek.Exps[0].String, &cronparse.WeekValue{Value: strings.ToLower(t)})
	}
}

func TestDayOfWeekStringRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * SUN-SAT *")
	assert.NoError(err)
	assert.Equal(cron.DayOfWeek.Exps[0].StringRange, &cronparse.WeekRange{From: "SUN", To: "SAT"})
}

func TestDayOfWeekComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * *,1,1-7,1/5,*/5,?,L,SUN,SUN-SAT *")
	assert.NoError(err)
	assert.Equal(cron.DayOfWeek.Exps[0].All, &cronparse.All{})
	assert.Equal(cron.DayOfWeek.Exps[1].Number, &cronparse.Number{Value: 1})
	assert.Equal(cron.DayOfWeek.Exps[2].NumberRange, &cronparse.NumberRange{From: 1, To: 7})
	assert.Equal(cron.DayOfWeek.Exps[3].Increment, &cronparse.Increment{Top: 1, Buttom: 5})
	assert.Equal(cron.DayOfWeek.Exps[4].Increment, &cronparse.Increment{Wildcard: true, Buttom: 5})
	assert.Equal(cron.DayOfWeek.Exps[5].Any, &cronparse.Any{})
	assert.Equal(cron.DayOfWeek.Exps[6].Last, &cronparse.Last{})
	assert.Equal(cron.DayOfWeek.Exps[7].String, &cronparse.WeekValue{Value: "SUN"})
	assert.Equal(cron.DayOfWeek.Exps[8].StringRange, &cronparse.WeekRange{From: "SUN", To: "SAT"})
}
