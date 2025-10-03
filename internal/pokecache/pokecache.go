package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	interval time.Duration
	entries  map[string]cacheEntry
	mu       *sync.RWMutex
}

func NewCache(duration time.Duration) Cache {
	mu := &sync.RWMutex{}
	cache := Cache{
		interval: duration,
		mu:       mu,
		entries:  make(map[string]cacheEntry),
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		val:       value,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(time.Now())
	}
}

func (c *Cache) reap(now time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		if entry.createdAt.Before(now.Add(-c.interval)) {
			delete(c.entries, key)
		}
	}
}
