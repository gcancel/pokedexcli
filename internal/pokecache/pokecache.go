package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	CacheEntries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	cache := new(Cache)
	cache.CacheEntries = make(map[string]cacheEntry, 0)
	select {
	case t := <-ticker.C:
		cache.reapLoop(t, interval)
	default:
		return cache
	}
	return cache
}

func (c *Cache) Add(k string, v []byte) {
	c.Lock()
	// copying underlying struct in map to be modified
	cache := c.CacheEntries[k]
	cache.createdAt = time.Now()
	cache.val = v
	// modifying the underlying struct with the new entry
	c.CacheEntries[k] = cache
	//fmt.Printf("from cache: %v \n", cache)
	c.Unlock()

}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.RLock()
	defer c.RUnlock()
	cache, exists := c.CacheEntries[k]
	if exists {
		return cache.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(t time.Time, i time.Duration) {
	//still need to make time.Ticker to make this function correctly...
	c.Lock()
	for key, cache := range c.CacheEntries {
		currentTime := t
		elapsed := currentTime.Add(-i)
		fmt.Println(elapsed)
		if elapsed.Sub(cache.createdAt) >= i {
			delete(c.CacheEntries, key)
		}
	}
	c.Unlock()
}
