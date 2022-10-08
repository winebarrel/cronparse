package cronparse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronparse"
)

func TestHoursAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? *")
	assert.NoError(err)
	assert.Equal(cron.Hours.Exps[0].All, &cronparse.All{})
}

func TestHoursNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* 0 * * ? *")
	assert.NoError(err)
	assert.Equal(cron.Hours.Exps[0].Number, &cronparse.Number{Value: 0})
}

func TestHoursNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* 0-23 * * ? *")
	assert.NoError(err)
	assert.Equal(cron.Hours.Exps[0].NumberRange, &cronparse.NumberRange{From: 0, To: 23})
}

func TestHoursIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* 0/5 * * ? *")
	assert.NoError(err)
	assert.Equal(cron.Hours.Exps[0].Increment, &cronparse.Increment{Top: 0, Buttom: 5})
}

func TestHoursIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* */5 * * ? *")
	assert.NoError(err)
	assert.Equal(cron.Hours.Exps[0].Increment, &cronparse.Increment{Wildcard: true, Buttom: 5})
}

func TestHoursComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* *,0,0-23,0/5,*/5 * * ? *")
	assert.NoError(err)
	assert.Equal(cron.Hours.Exps[0].All, &cronparse.All{})
	assert.Equal(cron.Hours.Exps[1].Number, &cronparse.Number{Value: 0})
	assert.Equal(cron.Hours.Exps[2].NumberRange, &cronparse.NumberRange{From: 0, To: 23})
	assert.Equal(cron.Hours.Exps[3].Increment, &cronparse.Increment{Top: 0, Buttom: 5})
	assert.Equal(cron.Hours.Exps[4].Increment, &cronparse.Increment{Wildcard: true, Buttom: 5})
}
