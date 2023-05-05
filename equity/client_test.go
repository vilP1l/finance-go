package equity

import (
	"testing"

	tests "github.com/piquette/finance-go/testing"
	"github.com/stretchr/testify/assert"
	finance "github.com/vildzi/finance-go"
)

func TestGetRegularMarketEquity(t *testing.T) {
	tests.SetMarket(finance.MarketStateRegular)

	q, err := Get(tests.TestEquitySymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, string(finance.MarketStateRegular), string(q.MarketState))
	assert.Equal(t, tests.TestEquitySymbol, q.Symbol)
}

func TestGetPostMarketEquity(t *testing.T) {
	tests.SetMarket(finance.MarketStatePost)

	q, err := Get(tests.TestEquitySymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, string(finance.MarketStatePost), string(q.MarketState))
	assert.Equal(t, tests.TestEquitySymbol, q.Symbol)
}

func TestGetPreMarketEquity(t *testing.T) {
	tests.SetMarket(finance.MarketStatePre)

	q, err := Get(tests.TestEquitySymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, string(finance.MarketStatePre), string(q.MarketState))
	assert.Equal(t, tests.TestEquitySymbol, q.Symbol)
}

func TestNilParamsEquity(t *testing.T) {

	iter := List(nil)

	assert.False(t, iter.Next())
	assert.Equal(t, "code: api-error, detail: missing function argument", iter.Err().Error())
}

func TestGetBadEquity(t *testing.T) {
	tests.SetMarket(finance.MarketStateRegular)

	q, err := Get("TEST")
	assert.Nil(t, q)
	assert.Nil(t, err)
}
