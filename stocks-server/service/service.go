package service

import (
	"fmt"
	"log"
	"sync"

	"rtp-stock-dashboard/stocks-server/model"
	"rtp-stock-dashboard/stocks-server/stockapi"
)

// StockService manages stock data
type StockService struct {
	stocks    map[string]*model.StockBoard
	watchlist map[string][]string // UserID -> List of stock symbols
	mu        sync.RWMutex
}

// NewStockService initializes the stock service
func NewStockService() *StockService {
	return &StockService{
		stocks:    make(map[string]*model.StockBoard),
		watchlist: make(map[string][]string),
	}
}

// GetStock fetches real-time stock data from Yahoo Finance API
func (s *StockService) GetStock(symbol string) (*model.StockBoard, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	stock, err := stockapi.FetchStockBoard(symbol)
	if err != nil {
		return nil, err
	}

	// Cache the fetched stock
	s.stocks[symbol] = stock

	return stock, nil
}

// UpdateStockPrice fetches and updates stock price from Yahoo Finance API
func (s *StockService) UpdateStockPrice(symbol string) (*model.StockBoard, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	stock, err := stockapi.FetchStockBoard(symbol)
	if err != nil {
		return nil, err
	}

	// Store updated stock data in service cache
	s.stocks[symbol] = stock
	log.Printf("Updated stock: %s - $%.2f", symbol, stock.Price)

	return stock, nil
}

// AddToWatchlist adds a stock to the user's watchlist
func (s *StockService) AddToWatchlist(userID, symbol string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.watchlist[userID] = append(s.watchlist[userID], symbol)
	log.Printf("User %s added %s to watchlist", userID, symbol)
	return true, nil
}

// RemoveFromWatchlist removes a stock from the user's watchlist
func (s *StockService) RemoveFromWatchlist(userID, symbol string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if stocks, exists := s.watchlist[userID]; exists {
		for i, sb := range stocks {
			if sb == symbol {
				s.watchlist[userID] = append(stocks[:i], stocks[i+1:]...)
				log.Printf("User %s removed %s from watchlist", userID, symbol)
				return true, nil
			}
		}
	}

	return false, fmt.Errorf("stock %s not found in watchlist for user %s", symbol, userID)
}
