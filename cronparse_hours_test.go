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
	assert.Equal(&cronparse.All{}, cron.Hours.Exps[0].All)
}

func TestHoursNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* 0 * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Number{Value: 0}, cron.Hours.Exps[0].Number)
}

func TestHoursNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* 0-23 * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.NumberRange{From: 0, To: 23}, cron.Hours.Exps[0].NumberRange)
}

func TestHoursIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* 0/5 * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Top: 0, Buttom: 5}, cron.Hours.Exps[0].Increment)
}

func TestHoursIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* */5 * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 5}, cron.Hours.Exps[0].Increment)
}

func TestHoursComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* *,0,0-23,0/5,*/5 * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.All{}, cron.Hours.Exps[0].All)
	assert.Equal(&cronparse.Number{Value: 0}, cron.Hours.Exps[1].Number)
	assert.Equal(&cronparse.NumberRange{From: 0, To: 23}, cron.Hours.Exps[2].NumberRange)
	assert.Equal(&cronparse.Increment{Top: 0, Buttom: 5}, cron.Hours.Exps[3].Increment)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 5}, cron.Hours.Exps[4].Increment)
}
