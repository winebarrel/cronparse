package cronparse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronparse"
)

func TestYearAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronparse.All{}, cron.Year.Exps[0].All)
}

func TestYearNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? 2022")
	assert.NoError(err)
	assert.Equal(&cronparse.Number{Value: 2022}, cron.Year.Exps[0].Number)
}

func TestYearNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? 1970-2199")
	assert.NoError(err)
	assert.Equal(&cronparse.NumberRange{From: 1970, To: 2199}, cron.Year.Exps[0].NumberRange)
}

func TestYearIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? 1970/2")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Top: 1970, Buttom: 2}, cron.Year.Exps[0].Increment)
}

func TestYearIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? */2")
	assert.NoError(err)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 2}, cron.Year.Exps[0].Increment)
}

func TestYearComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? *,2022,1970-2199,1970/2,*/2")
	assert.NoError(err)
	assert.Equal(&cronparse.All{}, cron.Year.Exps[0].All)
	assert.Equal(&cronparse.Number{Value: 2022}, cron.Year.Exps[1].Number)
	assert.Equal(&cronparse.NumberRange{From: 1970, To: 2199}, cron.Year.Exps[2].NumberRange)
	assert.Equal(&cronparse.Increment{Top: 1970, Buttom: 2}, cron.Year.Exps[3].Increment)
	assert.Equal(&cronparse.Increment{Wildcard: true, Buttom: 2}, cron.Year.Exps[4].Increment)
}
