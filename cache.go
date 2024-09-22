package main

import (
	"sync"
	"time"
)

type CacheItem struct {
	Key       string
	Value     string
	ExpiresAt time.Time
}

type LRUCache struct {
	capacity int
	cache    map[string]*CacheItem
	order    *DoublyLinkedList
	mutex    sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*CacheItem),
		order:    NewDoublyLinkedList(),
	}
}

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

func (c *LRUCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, exists := c.cache[key]; exists {
		c.order.Remove(item)
		delete(c.cache, key)
	}
}
