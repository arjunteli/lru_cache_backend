package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"lru_cache_backend/internal/cache"
)

// Handler manages the HTTP handlers for the cache operations
type Handler struct {
	cache *cache.LRUCache
}

// NewHandler creates a new Handler with the provided cache
func NewHandler(cache *cache.LRUCache) *Handler {
	return &Handler{cache: cache}
}

// GetHandler handles GET requests to retrieve a value from the cache
func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Processing GET request")

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is required", http.StatusBadRequest)
		return
	}

	value, exists := h.cache.Get(key)
	if !exists {
		fmt.Print("Key Not Found: ", key)
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

// SetHandler handles POST requests to add or update a value in the cache
func (h *Handler) SetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Processing SET request")

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	ttlStr := r.URL.Query().Get("ttl")

	if key == "" || value == "" {
		http.Error(w, "Key and value parameters are required", http.StatusBadRequest)
		return
	}

	ttl, err := strconv.Atoi(ttlStr)
	if err != nil || ttl <= 0 {
		fmt.Println("Invalid TTL: ", ttlStr, " Error: ", err)
		http.Error(w, "Invalid or missing TTL parameter", http.StatusBadRequest)
		return
	}

	fmt.Printf("Key=%s, Value=%s, TTL=%d\n", key, value, ttl)
	h.cache.Set(key, value, time.Duration(ttl)*time.Second)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
