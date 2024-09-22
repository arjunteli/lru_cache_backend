package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside getHandler")

	key := r.URL.Query().Get("key")
	value, exists := cache.Get(key)
	if !exists {
		fmt.Print("Key Not Found")
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
	if err != nil {
		return
	}
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside setHandler")
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	ttl, err := strconv.Atoi(r.URL.Query().Get("ttl"))

	fmt.Print(" Key=", key, ", Value=", value, ", TTL=", ttl)
	if err != nil {
		fmt.Println("Error ", err)
		http.Error(w, "Invalid TTL", http.StatusBadRequest)
		return
	}
	cache.Set(key, value, time.Duration(ttl)*time.Second)
	err = json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
	if err != nil {
		fmt.Println("Error ", err)
		return
	}

}
