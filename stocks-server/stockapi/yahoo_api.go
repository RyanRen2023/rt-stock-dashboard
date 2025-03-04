package stockapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rtp-stock-dashboard/stocks-server/model"
	"time"
)

// YahooFinanceResponse represents the JSON structure received from Yahoo Finance API.
type YahooFinanceResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Symbol        string  `json:"symbol"`
				Currency      string  `json:"currency"`
				Exchange      string  `json:"exchangeName"`
				MarketPrice   float64 `json:"regularMarketPrice"`
				PreviousClose float64 `json:"previousClose"`
				MarketTime    int64   `json:"regularMarketTime"`
			} `json:"meta"`
			Indicators struct {
				Quote []struct {
					Volume []int `json:"volume"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

// FetchStockBoard retrieves stock data from Yahoo Finance API and maps it to StockBoard struct.
func FetchStockBoard(symbol string) (*model.StockBoard, error) {

	yURL := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s", symbol)

	request, reqErr := http.NewRequest(http.MethodGet, yURL, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")

	response, reqErr := http.DefaultClient.Do(request)
	if reqErr != nil {
		return nil, reqErr
	}

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("received status code %d", response.StatusCode)
	}

	// Parse JSON response
	var yfResp YahooFinanceResponse
	if decErr := json.NewDecoder(response.Body).Decode(&yfResp); decErr != nil {
		return nil, decErr
	}

	// Validate response data
	if len(yfResp.Chart.Result) == 0 {
		return nil, fmt.Errorf("no stock data found for symbol: %s", symbol)
	}

	meta := yfResp.Chart.Result[0].Meta
	quote := yfResp.Chart.Result[0].Indicators.Quote

	// Ensure that volume data exists
	var volume int
	if len(quote) > 0 && len(quote[0].Volume) > 0 {
		volume = quote[0].Volume[0]
	} else {
		volume = 0 // Default value if volume data is missing
	}

	// Map Yahoo Finance data to StockBoard struct
	stockBoard := &model.StockBoard{
		Symbol:        meta.Symbol,
		Name:          fmt.Sprintf("%s Inc.", meta.Symbol), // Placeholder for real company name
		Price:         meta.MarketPrice,
		Change:        meta.MarketPrice - meta.PreviousClose,
		PercentChange: ((meta.MarketPrice - meta.PreviousClose) / meta.PreviousClose) * 100,
		Volume:        volume,
		Currency:      meta.Currency,
		Exchange:      meta.Exchange,
		UpdatedAt:     time.Unix(meta.MarketTime, 0), // Convert timestamp to time.Time
	}

	return stockBoard, nil
}
