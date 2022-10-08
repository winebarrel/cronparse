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
	assert.Equal(cron.Year.Exps[0].All, &cronparse.All{})
}

func TestYearNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? 2022")
	assert.NoError(err)
	assert.Equal(cron.Year.Exps[0].Number, &cronparse.Number{Value: 2022})
}

func TestYearNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? 1970-2199")
	assert.NoError(err)
	assert.Equal(cron.Year.Exps[0].NumberRange, &cronparse.NumberRange{From: 1970, To: 2199})
}

func TestYearIncrement(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? 1970/2")
	assert.NoError(err)
	assert.Equal(cron.Year.Exps[0].Increment, &cronparse.Increment{Top: 1970, Buttom: 2})
}

func TestYearIncrementWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? */2")
	assert.NoError(err)
	assert.Equal(cron.Year.Exps[0].Increment, &cronparse.Increment{Wildcard: true, Buttom: 2})
}

func TestYearComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronparse.Parser.ParseString("", "* * * * ? *,2022,1970-2199,1970/2,*/2")
	assert.NoError(err)
	assert.Equal(cron.Year.Exps[0].All, &cronparse.All{})
	assert.Equal(cron.Year.Exps[1].Number, &cronparse.Number{Value: 2022})
	assert.Equal(cron.Year.Exps[2].NumberRange, &cronparse.NumberRange{From: 1970, To: 2199})
	assert.Equal(cron.Year.Exps[3].Increment, &cronparse.Increment{Top: 1970, Buttom: 2})
	assert.Equal(cron.Year.Exps[4].Increment, &cronparse.Increment{Wildcard: true, Buttom: 2})
}
