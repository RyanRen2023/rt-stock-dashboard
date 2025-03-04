package model

import (
	"fmt"
	"time"
)

// StockBoard represents a real-time stock dashboard entry
type StockBoard struct {
	Symbol        string    `json:"symbol"`         // Stock symbol (e.g., "AAPL")
	Name          string    `json:"name"`           // Company name
	Price         float64   `json:"price"`          // Current market price
	Change        float64   `json:"change"`         // Price change (current price - previous close price)
	PercentChange float64   `json:"percent_change"` // Percentage change
	Volume        int       `json:"volume"`         // Trading volume
	Currency      string    `json:"currency"`       // Currency unit (e.g., "USD")
	Exchange      string    `json:"exchange"`       // Exchange market (e.g., "NASDAQ")
	UpdatedAt     time.Time `json:"updated_at"`     // Last updated timestamp
}

// UpdatePrice updates the stock price and recalculates the change percentage.
func (s *StockBoard) UpdatePrice(newPrice float64) {
	oldPrice := s.Price
	s.Price = newPrice
	s.Change = s.Price - oldPrice
	if oldPrice != 0 {
		s.PercentChange = (s.Change / oldPrice) * 100
	} else {
		s.PercentChange = 0 // Avoid division by zero
	}
	s.UpdatedAt = time.Now()
}

// SetVolume updates the trading volume.
func (s *StockBoard) SetVolume(newVolume int) {
	s.Volume = newVolume
}

// FormatStockData returns a formatted string for displaying stock information.
func (s *StockBoard) FormatStockData() string {
	return fmt.Sprintf(
		"Stock: %s (%s) | Price: %.2f %s | Change: %.2f (%.2f%%) | Volume: %d | Exchange: %s | Updated: %s",
		s.Name, s.Symbol, s.Price, s.Currency, s.Change, s.PercentChange, s.Volume, s.Exchange, s.UpdatedAt.Format(time.RFC3339),
	)
}

// IsPositiveChange checks if the stock price has increased.
func (s *StockBoard) IsPositiveChange() bool {
	return s.Change > 0
}

// IsNegativeChange checks if the stock price has decreased.
func (s *StockBoard) IsNegativeChange() bool {
	return s.Change < 0
}

// TimeAgo returns a human-readable time difference (e.g., "5 minutes ago").
func (s *StockBoard) TimeAgo() string {
	duration := time.Since(s.UpdatedAt)
	switch {
	case duration.Minutes() < 1:
		return "Just now"
	case duration.Hours() < 1:
		return fmt.Sprintf("%.0f minutes ago", duration.Minutes())
	case duration.Hours() < 24:
		return fmt.Sprintf("%.0f hours ago", duration.Hours())
	default:
		return s.UpdatedAt.Format("2006-01-02 15:04:05") // Standard date format
	}
}
