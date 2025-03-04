package stock_client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GetStockPrice function
func TestGetStockPrice(t *testing.T) {

	// Call the function
	stock, err := GetStockPrice("AAPL")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, stock)

	println("Stock price:", stock.Price)

	assert.Equal(t, "AAPL", stock.Symbol)

}
