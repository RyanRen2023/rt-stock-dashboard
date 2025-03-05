# Stock Service


## Project Overview
The **Stocks Server** is a backend service that provides stock market data. It integrates with **Yahoo Finance API** to fetch stock information and supports **REST API** for data retrieval.  


## Technology Stack
- **Go** (`1.20+`)
- **REST API** (Provides HTTP endpoints for stock data)
- **Yahoo Finance API** (Fetches stock price data)
- **Docker** (Containerization support)

## Directory Structure
```
stock-service
├── Dockerfile            # Docker configuration for containerization
├── README.md             # Project documentation
├── go.mod                # Go module dependencies
├── go.sum                # Dependency lock file
├── model                 # Data models
│   └── stock.go          # Defines stock data structure
├── service               # Business logic layer
│   ├── handler.go        # Handles HTTP requests
│   └── service.go        # Business logic implementation
├── stock_server.go       # Main application entry point
├── stockapi              # External API integration
│   ├── yahoo_api.go      # Yahoo Finance API integration
│   └── yahoo_api_test.go # Unit tests for Yahoo API
└── testdata
    └── yahoo.json        # Sample API response from Yahoo Finance

## Installation & Running the Server

### 1. Install Go Dependencies
Run the following command to install necessary dependencies:
```sh
go mod tidy
```

### 2. Run the Server
Start the stock service using:
```sh
go run stock_server.go
```

## API Endpoints

### REST API

| Method | Endpoint               | Description                                  |
|--------|-------------------------|----------------------------------------------|
| GET    | `/stock-price`          | Fetch the current stock price               |
| POST   | `/update-stock-price`   | Update the stock price                      |
| POST   | `/add-watchlist`        | Add a stock to the watchlist                |
| POST   | `/remove-watchlist`     | Remove a stock from the watchlist           |

## Docker Deployment
To build and run the service in a container:
```sh
docker build -t stock-service .
docker run -p 8081:8081 stock-service
```

## License
This project is licensed under the **MIT License**.