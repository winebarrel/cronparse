package cronparse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronparse"
)

func TestMinutesAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.All{}, cron.Minutes.Exps[0].All)
}

func TestMinutesNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "0 * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Number{Value: 0}, cron.Minutes.Exps[0].Number)
}

func TestMinutesNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "0-59 * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.NumberRange{From: 0, To: 59}, cron.Minutes.Exps[0].NumberRange)
}

func TestMinutesIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "0/5 * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Top: 0, Buttom: 5}, cron.Minutes.Exps[0].Increment)
}

func TestMinutesIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "*/5 * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 5}, cron.Minutes.Exps[0].Increment)
}

func TestMinutesComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "*,0,0-59,0/5,*/5 * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.All{}, cron.Minutes.Exps[0].All)
	assert.Equal(&cronparse.Number{Value: 0}, cron.Minutes.Exps[1].Number)
	assert.Equal(&cronparse.NumberRange{From: 0, To: 59}, cron.Minutes.Exps[2].NumberRange)
	assert.Equal(&cronparse.Increment{Top: 0, Buttom: 5}, cron.Minutes.Exps[3].Increment)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 5}, cron.Minutes.Exps[4].Increment)
}
