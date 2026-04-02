package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entries map[string]cacheEntry
	interval time.Duration
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(i time.Duration) Cache {
	entries := make(map[string]cacheEntry)
	c := Cache{
		interval: i,
		entries: entries,
	}
	go c.reapLoop()
	return c
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.entries[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, found := c.entries[key]

	return entry.val, found
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		for key, entry := range c.entries {
				age := time.Now().Sub(entry.createdAt)
				if age > c.interval {
					c.mu.Lock()
					defer c.mu.Unlock()

					delete(c.entries, key)
				}
			} 
	}
	
}