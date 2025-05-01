package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries	 map[string]cacheEntry
	mux   	 *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val 	  []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		entries: map[string]cacheEntry{},
		mux: new(sync.Mutex),
	}

	go newCache.reapLoop(interval)
	return newCache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	cache.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	val := []byte{}

	cache.mux.Lock()
	defer cache.mux.Unlock()
	entry, exists := cache.entries[key]

	if exists {
		val = entry.val
	}
	return val, exists
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	
	go func() {
		for range ticker.C {
			cache.mux.Lock()
			for key, val := range cache.entries {
				if time.Since(val.createdAt) > interval {
					delete(cache.entries, key)
				} 
			}
			cache.mux.Unlock()
		}
	}()
}