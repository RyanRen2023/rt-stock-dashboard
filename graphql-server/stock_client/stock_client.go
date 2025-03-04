package stock_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"rpt-stock-dashboard/graphql-server/graph/model"
	"time"
)

// // Stock struct represents stock data
// type Stock struct {
// 	Symbol        string    `json:"symbol"`         // Stock symbol (e.g., "AAPL")
// 	Name          string    `json:"name"`           // Company name
// 	Price         float64   `json:"price"`          // Current market price
// 	Change        float64   `json:"change"`         // Price change (current price - previous close price)
// 	PercentChange float64   `json:"percent_change"` // Percentage change
// 	Volume        int32     `json:"volume"`         // Trading volume
// 	Currency      string    `json:"currency"`       // Currency unit (e.g., "USD")
// 	Exchange      string    `json:"exchange"`       // Exchange market (e.g., "NASDAQ")
// 	UpdatedAt     time.Time `json:"updated_at"`     // Last updated timestamp`
// }

// GetStockPrice simulates fetching the current stock price
func GetStockPrice(symbol string) (*model.Stock, error) {

	// Read environment variable API_URL, default to http://localhost:8081
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8081" // Default value
	}

	// Construct the URL
	url := fmt.Sprintf("%s/stock-price?symbol=%s", apiURL, symbol)

	// Set up HTTP client with timeout
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("network request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch stock data: received status code %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Debugging: Print response body
	fmt.Println("Response Body:", string(body))

	// Parse JSON data using graph.model.Stock
	var stock *model.Stock
	if err := json.Unmarshal(body, &stock); err != nil {
		return nil, fmt.Errorf("failed to parse stock data: %v", err)
	}

	// Ensure UpdatedAt is in the correct time format
	stock.UpdatedAt = time.Now().Format(time.RFC3339)

	return stock, nil

}
