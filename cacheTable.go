package mycache

import (
	"sync"
	"time"
)

type CacheTable struct {
	name string
	sync.RWMutex
	items map[interface{}]*CacheItem
}

func (table *CacheTable) Count() int {
	table.RLock()
	defer table.RUnlock()
	return len(table.items)
}

func (table *CacheTable) Add(key, value interface{}, lifeSpan time.Duration) *CacheItem {
	item := newItem(key, value, lifeSpan)
	table.Lock()
	table.items[key] = item
	table.Unlock()
	return item
}

func (table *CacheTable) AddIfNotExists(key, value interface{}, lifeSpan time.Duration) (*CacheItem, bool) {
	table.RLock()
	c, ok := table.items[key]
	table.RUnlock()
	if ok {
		return c, false
	}
	return table.Add(key, value, lifeSpan), true
}

func (table *CacheTable) Delete(key interface{}) (*CacheItem, error) {
	table.Lock()
	v, ok := table.items[key]
	defer table.Unlock()
	if !ok {
		return nil, ErrKeyNotFound(key)
	}
	delete(table.items, key)
	return v, nil
}

func (table *CacheTable) Exists(key interface{}) bool {
	table.RLock()
	defer table.RUnlock()
	_, ok := table.items[key]
	return ok
}

func (table *CacheTable) Get(key interface{}) (interface{}, error) {
	table.RLock()
	v, ok := table.items[key]
	table.RUnlock()
	if !ok {
		return nil, ErrKeyNotFound(key)
	}
	return v, nil
}
