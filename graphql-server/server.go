package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"rpt-stock-dashboard/graphql-server/graph"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := &graph.Resolver{}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{ // Add WebSocket subscription support
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Allow all origins to access
			},
		},
	})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all domains to access
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	go func() {
		fmt.Println("Start sending stock updates")
		for {
			time.Sleep(5 * time.Second)

			// Get all subscribed stock codes
			stockCodes := resolver.Observers()

			// Iterate through subscribed stock codes and fetch the latest data using GraphQL Query
			for _, stockCode := range stockCodes {
				// Call GraphQL Query instead of directly calling `stock_client.GetStockPrice`
				stock, err := resolver.Query().Stock(context.Background(), stockCode)
				if err != nil {
					fmt.Printf("Failed to get stock data for %s: %v\n", stockCode, err)
					continue
				}
				// Push update to subscribers
				resolver.PublishStockPrice(stock)
				fmt.Printf("Sent stock update: %s - $%.2f\n", stock.Symbol, stock.Price)
			}
		}
	}()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
