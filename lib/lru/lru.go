package lru

import (
	"gofly/lib/lru/simplelru"
	"sync"
)

var cacheMap = make(map[string]*Cache, 0)
var lockForCacheMap sync.RWMutex

// Cache is a thread-safe fixed size LRU cache.
type Cache struct {
	lru  simplelru.LRUCache
	lock sync.RWMutex
}

// New creates an LRU of the given size.
func New(size int) (*Cache, error) {
	return NewWithEvict(size, nil)
}

// NewWithEvict constructs a fixed size cache with the given eviction
// callback.
func NewWithEvict(size int, onEvicted func(key interface{}, value interface{})) (*Cache, error) {
	lru, err := simplelru.NewLRU(size, simplelru.EvictCallback(onEvicted))
	if err != nil {
		return nil, err
	}
	c := &Cache{
		lru: lru,
	}
	return c, nil
}

func New2(key string, size int) (*Cache, error) {
	v := Get(key)
	if v != nil {
		return v, nil
	}

	lockForCacheMap.Lock()
	defer lockForCacheMap.Unlock()

	v, err := New(size)
	if err != nil {
		return nil, err
	}

	cacheMap[key] = v

	return v, nil
}

// Get 通过 key 获取 cacheMap 里的对象
func Get(key string) *Cache {
	return cacheMap[key]
}

// GetMultiLruStatus 批量获取 LRU 统计信息。不传则返回 `cacheMap` 中的全部。
func GetMultiLruStatus(keys ...string) map[string]*simplelru.LRUStat {
	var res = make(map[string]*simplelru.LRUStat)
	if len(keys) == 0 {
		for key, cache := range cacheMap {
			res[key] = cache.Stats()
		}
		return res
	}

	for _, key := range keys {
		if v, has := cacheMap[key]; has {
			res[key] = v.Stats()
		}
	}

	return res
}

// GetLruStatus 获取单个 lru cache 的 统计信息
func GetLruStatus(key string) *simplelru.LRUStat {
	return GetMultiLruStatus(key)[key]
}

// Purge is used to completely clear the cache.
func (c *Cache) Purge() {
	c.lock.Lock()
	c.lru.Purge()
	c.lock.Unlock()
}

// get stats
func (c *Cache) Stats() *simplelru.LRUStat {
	return c.lru.Stats()
}

// Add adds a value to the cache.  Returns true if an eviction occurred.
func (c *Cache) Set(key, value interface{}, expiresIn int64) (evicted bool) {
	c.lock.Lock()
	evicted = c.lru.Set(key, value, expiresIn)
	c.lock.Unlock()
	return evicted
}

// Get looks up a key's value from the cache.
func (c *Cache) Get(key interface{}) (value interface{}, ok bool) {
	c.lock.Lock()
	value, ok = c.lru.Get(key)
	c.lock.Unlock()
	return value, ok
}

// Contains checks if a key is in the cache, without updating the
// recent-ness or deleting it for being stale.
func (c *Cache) Contains(key interface{}) bool {
	c.lock.RLock()
	containKey := c.lru.Contains(key)
	c.lock.RUnlock()
	return containKey
}

// Peek returns the key value (or undefined if not found) without updating
// the "recently used"-ness of the key.
func (c *Cache) Peek(key interface{}) (value interface{}, ok bool) {
	c.lock.RLock()
	value, ok = c.lru.Peek(key)
	c.lock.RUnlock()
	return value, ok
}

// Cas checks if a key is in the cache  without updating the
// recent-ness or deleting it for being stale,  and if not, adds the value.
// Returns whether found and whether an eviction occurred.
func (c *Cache) Cas(key, value interface{}, expiresIn int64) (ok, evicted bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.lru.Contains(key) {
		return true, false
	}
	evicted = c.lru.Set(key, value, expiresIn)
	return false, evicted
}

// Remove removes the provided key from the cache.
func (c *Cache) Remove(key interface{}) (present bool) {
	c.lock.Lock()
	present = c.lru.Remove(key)
	c.lock.Unlock()
	return
}

// Resize changes the cache size.
func (c *Cache) Resize(size int) (evicted int) {
	c.lock.Lock()
	evicted = c.lru.Resize(size)
	c.lock.Unlock()
	return evicted
}

// RemoveOldest removes the oldest item from the cache.
func (c *Cache) RemoveOldest() (key interface{}, value interface{}, ok bool) {
	c.lock.Lock()
	key, value, ok = c.lru.RemoveOldest()
	c.lock.Unlock()
	return
}

// GetOldest returns the oldest entry
func (c *Cache) GetOldest() (key interface{}, value interface{}, ok bool) {
	c.lock.Lock()
	key, value, ok = c.lru.GetOldest()
	c.lock.Unlock()
	return
}

// Keys returns a slice of the keys in the cache, from oldest to newest.
func (c *Cache) Keys() []interface{} {
	c.lock.RLock()
	keys := c.lru.Keys()
	c.lock.RUnlock()
	return keys
}

// Len returns the number of items in the cache.
func (c *Cache) Len() int {
	c.lock.RLock()
	length := c.lru.Len()
	c.lock.RUnlock()
	return length
}
