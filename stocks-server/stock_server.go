package main

import (
	"log"
	"net/http"
	"rtp-stock-dashboard/stocks-server/service"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize services
	stockService := service.NewStockService()
	stockHandler := service.NewStockHandler(stockService)

	// Setup routes
	r := mux.NewRouter()
	r.HandleFunc("/stock-price", stockHandler.GetStockPrice).Methods("GET")
	r.HandleFunc("/update-stock-price", stockHandler.UpdateStockPrice).Methods("POST")
	r.HandleFunc("/add-watchlist", stockHandler.AddToWatchlist).Methods("POST")
	r.HandleFunc("/remove-watchlist", stockHandler.RemoveFromWatchlist).Methods("POST")

	// Start server
	port := "8081"
	log.Printf("Stock server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
