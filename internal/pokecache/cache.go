package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mux *sync.Mutex
	cache map[string]cacheEntry
}
func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{createdAt: time.Now(), val: val}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cache, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return cache.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache {
		mux: &sync.Mutex{},
		cache: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}