package chart

import (
	"testing"

	"github.com/stretchr/testify/assert"
	tests "github.com/vildzi/finance-go/testing"
)

func TestGetEquityChart(t *testing.T) {
	p := &Params{Symbol: tests.TestEquitySymbol}
	iter := Get(p)
	assert.True(t, iter.Next())
}

func TestGetETFChart(t *testing.T) {
	p := &Params{Symbol: tests.TestETFSymbol}
	iter := Get(p)
	assert.True(t, iter.Next())
}

func TestGetFutureChart(t *testing.T) {
	p := &Params{Symbol: tests.TestFutureSymbol}
	iter := Get(p)
	assert.True(t, iter.Next())
}

func TestGetIndexChart(t *testing.T) {
	p := &Params{Symbol: tests.TestIndexSymbol}
	iter := Get(p)
	assert.True(t, iter.Next())
}

func TestGetOptionChart(t *testing.T) {
	p := &Params{Symbol: tests.TestOptionSymbol}
	chart := Get(p)
	assert.True(t, chart.Next())
}

func TestGetMutualFundChart(t *testing.T) {
	p := &Params{Symbol: tests.TestMutualFundSymbol}
	chart := Get(p)
	assert.True(t, chart.Next())
}

func TestGetForexPairChart(t *testing.T) {
	p := &Params{Symbol: tests.TestForexPairSymbol}
	chart := Get(p)
	assert.True(t, chart.Next())
}

func TestBadSymbolChart(t *testing.T) {
	p := &Params{Symbol: "BADSYMBOL"}
	chart := Get(p)
	assert.False(t, chart.Next())
	assert.NotNil(t, chart.Err())
}
