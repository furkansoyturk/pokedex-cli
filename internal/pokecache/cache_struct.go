package pokecache

import (
	"sync"
	"time"
)

type CacheStruct struct {
	entries map[string]cacheEntryStruct
	mux     *sync.Mutex
}
type cacheEntryStruct struct {
	createdAt time.Time
	val       []byte
}

func NewCache(timeToLive time.Duration) CacheStruct {
	cache := CacheStruct{
		entries: make(map[string]cacheEntryStruct),
		mux:     &sync.Mutex{},
	}
	go cache.CleanExpired(timeToLive.Seconds())
	return cache
}

func (cache *CacheStruct) Add(key string, val []byte) {
	cache.mux.Lock()
	defer cache.mux.Unlock()

	cache.entries[key] = cacheEntryStruct{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (cache *CacheStruct) Get(key string) ([]byte, bool) {
	entry, exist := cache.entries[key]
	return entry.val, exist
}

func (cache *CacheStruct) CleanExpired(timeTolive float64) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	for key := range cache.entries {
		createdAt := cache.entries[key].createdAt.Second()
		if float64(createdAt) + timeTolive >= float64(time.Now().UTC().Second()) {
			delete(cache.entries, key)
		}
	}
}
