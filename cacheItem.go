package mycache

import (
	"time"
)

type CacheItem struct {
	key         interface{}
	value       interface{}
	lifeSpan    time.Duration
	lastAccess  time.Time
	accessCount int
}

func newItem(key, value interface{}, lifeSpan time.Duration) *CacheItem {
	return &CacheItem{
		key:         key,
		value:       value,
		lifeSpan:    lifeSpan,
		lastAccess:  time.Now(),
		accessCount: 0,
	}
}
