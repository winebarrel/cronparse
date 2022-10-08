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
	assert.Equal(cron.Month.Exps[0].All, &cronparse.All{})
}

func TestMonthNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * 1 ? *")
	assert.NoError(err)
	assert.Equal(cron.Month.Exps[0].Number, &cronparse.Number{Value: 1})
}

func TestMonthNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * 1-12 ? *")
	assert.NoError(err)
	assert.Equal(cron.Month.Exps[0].NumberRange, &cronparse.NumberRange{From: 1, To: 12})
}

func TestMonthIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * 1/5 ? *")
	assert.NoError(err)
	assert.Equal(cron.Month.Exps[0].Increment, &cronparse.Increment{Top: 1, Buttom: 5})
}

func TestMonthIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * */5 ? *")
	assert.NoError(err)
	assert.Equal(cron.Month.Exps[0].Increment, &cronparse.Increment{Wildcard: true, Buttom: 5})
}

func TestMonthString(t *testing.T) {
	assert := assert.New(t)
	tt := []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

	for _, t := range tt {
		cron, err := cronparse.Parser.ParseString("", fmt.Sprintf("* * * %s ? *", t))
		assert.NoError(err)
		assert.Equal(cron.Month.Exps[0].String, &cronparse.MonthValue{Value: t})

		cron, err = cronparse.Parser.ParseString("", fmt.Sprintf("* * * %s ? *", strings.ToLower(t)))
		assert.NoError(err)
		assert.Equal(cron.Month.Exps[0].String, &cronparse.MonthValue{Value: strings.ToLower(t)})
	}
}

func TestMonthStringRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * JAN-DEC ? *")
	assert.NoError(err)
	assert.Equal(cron.Month.Exps[0].StringRange, &cronparse.MonthRange{From: "JAN", To: "DEC"})
}

func TestMonthComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * *,1,1-12,1/5,*/5,JAN,JAN-DEC ? *")
	assert.NoError(err)
	assert.Equal(cron.Month.Exps[0].All, &cronparse.All{})
	assert.Equal(cron.Month.Exps[1].Number, &cronparse.Number{Value: 1})
	assert.Equal(cron.Month.Exps[2].NumberRange, &cronparse.NumberRange{From: 1, To: 12})
	assert.Equal(cron.Month.Exps[3].Increment, &cronparse.Increment{Top: 1, Buttom: 5})
	assert.Equal(cron.Month.Exps[4].Increment, &cronparse.Increment{Wildcard: true, Buttom: 5})
	assert.Equal(cron.Month.Exps[5].String, &cronparse.MonthValue{Value: "JAN"})
	assert.Equal(cron.Month.Exps[6].StringRange, &cronparse.MonthRange{From: "JAN", To: "DEC"})
}
