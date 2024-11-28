package cache

import (
	"fmt"
	"sync"
	"time"
)

// Кэш на мапе с использованием RWMutex для безопасноти
type Cache struct {
	cache map[string]cacheItem
	mu    sync.RWMutex
}

// Элемент кэша, который содержит значение и exp. time
type cacheItem struct {
	value      interface{}
	expiration time.Time
}

// Конструктор
func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]cacheItem),
	}
}

// Set добавляет значение в кэш с ключем и TTL
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheItem{
		value:      value,
		expiration: time.Now().Add(ttl),
	}
}

// Get возвращает значение из кэша по ключу, если оно живое
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if item, found := c.cache[key]; found {
		if time.Now().Before(item.expiration) {
			return item.value, true
		}
		delete(c.cache, key)
	}
	return nil, false
}

// Delete удаляет значение из кэша по клчюу

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	c.mu.Unlock()

	delete(c.cache, key)
}

func (c *Cache) PrintCache() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for key, item := range c.cache {
		fmt.Printf("Ключ:%s, Value:%v, ttl: %s\n", key, item.value, item.expiration)
	}
}
