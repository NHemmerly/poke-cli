package pokecache

import (
	"maps"
	"sync"
	"time"
)

type (
	Cache struct {
		Content map[string]CacheEntry
		mu      sync.Mutex
	}
	CacheEntry struct {
		createdAt time.Time
		val       []byte
	}
)

type CacheInterface interface {
	Add(key string, val []byte)
	Get(key string) ([]byte, bool)
}

func (c *Cache) Add(key string, val []byte) {
	c.Content[key] = CacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if val, ok := c.Content[key]; ok {
		return val.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key := range maps.Keys(c.Content) {
			if time.Since(c.Content[key].createdAt) > interval {
				delete(c.Content, key)
			}
		}
		c.mu.Unlock()

	}

}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{Content: make(map[string]CacheEntry)}
	go cache.reapLoop(interval)
	return cache
}
