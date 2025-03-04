package service

import (
	"encoding/json"
	"net/http"
)

// StockHandler manages HTTP requests for stock data
type StockHandler struct {
	service *StockService
}

// NewStockHandler creates a new StockHandler
func NewStockHandler(service *StockService) *StockHandler {
	return &StockHandler{service: service}
}

// GET /stock-price?symbol=AAPL
func (h *StockHandler) GetStockPrice(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		http.Error(w, "Missing symbol", http.StatusBadRequest)
		return
	}

	stock, err := h.service.GetStock(symbol)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(stock)
}

// POST /update-stock-price
func (h *StockHandler) UpdateStockPrice(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Symbol string  `json:"symbol"`
		Price  float64 `json:"price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	stock, err := h.service.UpdateStockPrice(request.Symbol)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stock)
}

// POST /add-watchlist
func (h *StockHandler) AddToWatchlist(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UserID string `json:"userId"`
		Symbol string `json:"symbol"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	success, err := h.service.AddToWatchlist(request.UserID, request.Symbol)
	if err != nil || !success {
		http.Error(w, "Failed to add to watchlist", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// POST /remove-watchlist
func (h *StockHandler) RemoveFromWatchlist(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UserID string `json:"userId"`
		Symbol string `json:"symbol"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	success, err := h.service.RemoveFromWatchlist(request.UserID, request.Symbol)
	if err != nil || !success {
		http.Error(w, "Failed to remove from watchlist", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
