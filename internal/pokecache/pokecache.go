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
	mu       sync.Mutex
	entries  map[string]cacheEntry
	interval time.Duration
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, ok := cache.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()

	for range ticker.C {
		cache.mu.Lock()

		for key, entry := range cache.entries {
			if time.Since(entry.createdAt) > cache.interval {
				delete(cache.entries, key)
			}
		}
		cache.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	go cache.reapLoop()
	return cache
}
