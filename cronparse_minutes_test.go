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
	assert.Equal(cron.Minutes.Exps[0].All, &cronparse.All{})
}

func TestMinutesNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "0 * * * ? *")
	assert.NoError(err)
	assert.Equal(cron.Minutes.Exps[0].Number, &cronparse.Number{Value: 0})
}

func TestMinutesNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "0-59 * * * ? *")
	assert.NoError(err)
	assert.Equal(cron.Minutes.Exps[0].NumberRange, &cronparse.NumberRange{From: 0, To: 59})
}

func TestMinutesIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "0/5 * * * ? *")
	assert.NoError(err)
	assert.Equal(cron.Minutes.Exps[0].Increment, &cronparse.Increment{Top: 0, Buttom: 5})
}

func TestMinutesIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "*/5 * * * ? *")
	assert.NoError(err)
	assert.Equal(cron.Minutes.Exps[0].Increment, &cronparse.Increment{Wildcard: true, Buttom: 5})
}

func TestMinutesComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "*,0,0-59,0/5,*/5 * * * ? *")
	assert.NoError(err)
	assert.Equal(cron.Minutes.Exps[0].All, &cronparse.All{})
	assert.Equal(cron.Minutes.Exps[1].Number, &cronparse.Number{Value: 0})
	assert.Equal(cron.Minutes.Exps[2].NumberRange, &cronparse.NumberRange{From: 0, To: 59})
	assert.Equal(cron.Minutes.Exps[3].Increment, &cronparse.Increment{Top: 0, Buttom: 5})
	assert.Equal(cron.Minutes.Exps[4].Increment, &cronparse.Increment{Wildcard: true, Buttom: 5})
}
