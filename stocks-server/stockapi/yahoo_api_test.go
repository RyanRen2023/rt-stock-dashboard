package stockapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockYahooAPIResponse is a sample mock response from Yahoo Finance API
const mockYahooAPIResponse = `{
	"chart": {
		"result": [{
			"meta": {
				"symbol": "AAPL",
				"currency": "USD",
				"exchangeName": "NASDAQ",
				"regularMarketPrice": 150.50,
				"previousClose": 148.30,
				"regularMarketTime": 1740594351
			},
			"indicators": {
				"quote": [{
					"volume": [50000000]
				}]
			}
		}]
	}
}`

// TestFetchStockBoard_BasicCheck ensures the function returns a non-nil stock object
func TestFetchStockBoard_BasicCheck(t *testing.T) {
	// Create a mock server to simulate Yahoo Finance API response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockYahooAPIResponse))
	}))
	defer mockServer.Close()

	// Call FetchStockBoard with a test symbol
	stock, err := FetchStockBoard("AAPL")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Ensure stock is not nil
	if stock == nil {
		t.Fatalf("Expected stock data, got nil")
	}

	t.Logf("Stock data fetched successfully: %+v", stock)
}
