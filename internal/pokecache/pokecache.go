package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	CacheEntries map[string]CacheEntry
}

type CacheEntry struct {
	CreatedAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	//move ticker logic to repl, rework to just check the  expiration and return a new cache
	cache := Cache{CacheEntries: make(map[string]CacheEntry, 0)}
	ticker := time.NewTicker(interval)
	go func(){
		for{
			select{
			case <-ticker.C:
				currentTime := time.Now()
				cache.reapLoop(currentTime, interval)
			}
		}
		}()
	return &cache
}

func (c *Cache) Add(k string, v []byte) {
	c.Lock()
	defer c.Unlock()
	// copying underlying struct in map to be modified
	cache := c.CacheEntries[k]
	cache.CreatedAt = time.Now()
	cache.val = v
	// modifying the underlying struct with the new entry
	c.CacheEntries[k] = cache
	

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

func (c *Cache) reapLoop(t time.Time, interval time.Duration) {
	c.Lock()
	for key, cache := range c.CacheEntries {
		currentTime := t
		elapsed := currentTime.Add(-interval)
		if elapsed.Sub(cache.CreatedAt) >= interval {
			delete(c.CacheEntries, key)
			fmt.Printf("deleted: %v\n", key)
		}
	}
	c.Unlock()
}
