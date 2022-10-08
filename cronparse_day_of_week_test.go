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
	assert.Equal(&cronparse.All{}, cron.DayOfWeek.Exps[0].All)
}

func TestDayOfWeekNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * 1 *")
	assert.NoError(err)
	assert.Equal(&cronparse.Number{Value: 1}, cron.DayOfWeek.Exps[0].Number)
}

func TestDayOfWeekNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * 1-7 *")
	assert.NoError(err)
	assert.Equal(&cronparse.NumberRange{From: 1, To: 7}, cron.DayOfWeek.Exps[0].NumberRange)
}

func TestDayOfWeekIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * 1/5 *")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Top: 1, Buttom: 5}, cron.DayOfWeek.Exps[0].Increment)
}

func TestDayOfWeekIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * */5 *")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 5}, cron.DayOfWeek.Exps[0].Increment)
}

func TestDayOfWeekAny(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Any{}, cron.DayOfWeek.Exps[0].Any)
}

func TestDayOfWeekLast(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * L *")
	assert.NoError(err)
	assert.Equal(&cronparse.Last{}, cron.DayOfWeek.Exps[0].Last)
}

func TestDayOfWeekName(t *testing.T) {
	assert := assert.New(t)
	tt := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}

	for _, t := range tt {
		cron, err := cronparse.Parser.ParseString("", fmt.Sprintf("* * ? * %s *", t))
		assert.NoError(err)
		assert.Equal(&cronparse.WeekName{Value: t}, cron.DayOfWeek.Exps[0].Name, t)

		cron, err = cronparse.Parser.ParseString("", fmt.Sprintf("* * ? * %s *", strings.ToLower(t)))
		assert.NoError(err)
		assert.Equal(&cronparse.WeekName{Value: strings.ToLower(t)}, cron.DayOfWeek.Exps[0].Name, t)
	}
}

func TestDayOfWeekNameRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * SUN-SAT *")
	assert.NoError(err)
	assert.Equal(&cronparse.WeekRange{From: "SUN", To: "SAT"}, cron.DayOfWeek.Exps[0].NameRange)
}

func TestDayOfWeekComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * ? * *,1,1-7,1/5,*/5,?,L,SUN,SUN-SAT *")
	assert.NoError(err)
	assert.Equal(&cronparse.All{}, cron.DayOfWeek.Exps[0].All)
	assert.Equal(&cronparse.Number{Value: 1}, cron.DayOfWeek.Exps[1].Number)
	assert.Equal(&cronparse.NumberRange{From: 1, To: 7}, cron.DayOfWeek.Exps[2].NumberRange)
	assert.Equal(&cronparse.Increment{Top: 1, Buttom: 5}, cron.DayOfWeek.Exps[3].Increment)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 5}, cron.DayOfWeek.Exps[4].Increment)
	assert.Equal(&cronparse.Any{}, cron.DayOfWeek.Exps[5].Any)
	assert.Equal(&cronparse.Last{}, cron.DayOfWeek.Exps[6].Last)
	assert.Equal(&cronparse.WeekName{Value: "SUN"}, cron.DayOfWeek.Exps[7].Name)
	assert.Equal(&cronparse.WeekRange{From: "SUN", To: "SAT"}, cron.DayOfWeek.Exps[8].NameRange)
}
