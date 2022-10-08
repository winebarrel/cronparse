package cronparse_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronparse"
)

func TestMonthAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.All{}, cron.Month.Exps[0].All)
}

func TestMonthNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * 1 ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Number{Value: 1}, cron.Month.Exps[0].Number)
}

func TestMonthNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * 1-12 ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.NumberRange{From: 1, To: 12}, cron.Month.Exps[0].NumberRange)
}

func TestMonthIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * 1/5 ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Top: 1, Buttom: 5}, cron.Month.Exps[0].Increment)
}

func TestMonthIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * */5 ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 5}, cron.Month.Exps[0].Increment, t)
}

func TestMonthName(t *testing.T) {
	assert := assert.New(t)
	tt := []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

	for _, t := range tt {
		cron, err := cronparse.Parser.ParseString("", fmt.Sprintf("* * * %s ? *", t))
		assert.NoError(err)
		assert.Equal(&cronparse.MonthName{Value: t}, cron.Month.Exps[0].Name)

		cron, err = cronparse.Parser.ParseString("", fmt.Sprintf("* * * %s ? *", strings.ToLower(t)))
		assert.NoError(err)
		assert.Equal(&cronparse.MonthName{Value: strings.ToLower(t)}, cron.Month.Exps[0].Name, t)
	}
}

func TestMonthNameRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * JAN-DEC ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.MonthRange{From: "JAN", To: "DEC"}, cron.Month.Exps[0].NameRange)
}

func TestMonthComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * *,1,1-12,1/5,*/5,JAN,JAN-DEC ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.All{}, cron.Month.Exps[0].All)
	assert.Equal(&cronparse.Number{Value: 1}, cron.Month.Exps[1].Number)
	assert.Equal(&cronparse.NumberRange{From: 1, To: 12}, cron.Month.Exps[2].NumberRange)
	assert.Equal(&cronparse.Increment{Top: 1, Buttom: 5}, cron.Month.Exps[3].Increment)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 5}, cron.Month.Exps[4].Increment)
	assert.Equal(&cronparse.MonthName{Value: "JAN"}, cron.Month.Exps[5].Name)
	assert.Equal(&cronparse.MonthRange{From: "JAN", To: "DEC"}, cron.Month.Exps[6].NameRange)
}
