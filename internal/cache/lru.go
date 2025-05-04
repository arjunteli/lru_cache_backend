package cache

import (
	"sync"
	"time"
)

// CacheItem represents an item stored in the cache
type CacheItem struct {
	Key       string
	Value     string
	ExpiresAt time.Time
}

// LRUCache implements a Least Recently Used cache with TTL support
type LRUCache struct {
	capacity int
	cache    map[string]*CacheItem
	order    *DoublyLinkedList
	mutex    sync.RWMutex
}

// NewLRUCache creates a new LRU cache with the specified capacity
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*CacheItem),
		order:    NewDoublyLinkedList(),
	}
}

// Get retrieves a value from the cache by key
func (c *LRUCache) Get(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, exists := c.cache[key]
	if !exists || time.Now().After(item.ExpiresAt) {
		return "", false
	}
	c.order.MoveToFront(item)
	return item.Value, true
}

// Set adds or updates a key-value pair in the cache with a TTL
func (c *LRUCache) Set(key string, value string, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, exists := c.cache[key]; exists {
		item.Value = value
		item.ExpiresAt = time.Now().Add(ttl)
		c.order.MoveToFront(item)
		return
	}

	if len(c.cache) >= c.capacity {
		oldest := c.order.RemoveOldest()
		delete(c.cache, oldest.Key)
	}

	newItem := &CacheItem{
		Key:       key,
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}
	c.cache[key] = newItem
	c.order.AddToFront(newItem)
}

// Delete removes a key from the cache
func (c *LRUCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, exists := c.cache[key]; exists {
		c.order.Remove(item)
		delete(c.cache, key)
	}
}
