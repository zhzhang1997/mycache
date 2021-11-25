package mycache

import (
	"sync"
)

var (
	mutex sync.RWMutex
	cache = map[string]*CacheTable{}
)

func Cache(name string) *CacheTable {
	mutex.RLock()
	c, ok := cache[name]
	mutex.RUnlock()
	if !ok {
		mutex.Lock()
		c, ok = cache[name]
		if !ok {
			c = &CacheTable{
				name:  name,
				items: make(map[interface{}]*CacheItem),
			}
			cache[name] = c
		}
		mutex.Unlock()
	}
	return c
}
