package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"lru_cache_backend/internal/api"
	"lru_cache_backend/internal/cache"
)

func main() {
	// Get cache capacity from environment variable or use default
	capacityStr := os.Getenv("CACHE_CAPACITY")
	capacity := 1024 // Default capacity
	if capacityStr != "" {
		var err error
		capacity, err = strconv.Atoi(capacityStr)
		if err != nil {
			log.Fatalf("Invalid CACHE_CAPACITY value: %v", err)
		}
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Initialize cache
	lruCache := cache.NewLRUCache(capacity)

	// Initialize API handlers
	handler := api.NewHandler(lruCache)

	// Setup routes with middleware
	http.Handle("/get", api.CorsMiddleware(http.HandlerFunc(handler.GetHandler)))
	http.Handle("/set", api.CorsMiddleware(http.HandlerFunc(handler.SetHandler)))

	// Start server
	serverAddr := ":" + port
	fmt.Printf("Server starting on port %s with cache capacity %d\n", port, capacity)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
