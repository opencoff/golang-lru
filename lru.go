package lru

import (
	"sync"

	"github.com/opencoff/golang-lru/simplelru"
)

// SimpleCache is a thread-safe fixed size LRU cache.
type SimpleCache struct {
	lru  simplelru.LRUCache
	lock sync.RWMutex
}

// New creates an LRU of the given size.
func NewSimple(size int) (*SimpleCache, error) {
	return NewSimpleWithEvict(size, nil)
}

// NewWithEvict constructs a fixed size cache with the given eviction
// callback.
func NewSimpleWithEvict(size int, onEvicted func(key interface{}, value interface{})) (*SimpleCache, error) {
	lru, err := simplelru.NewLRU(size, simplelru.EvictCallback(onEvicted))
	if err != nil {
		return nil, err
	}
	c := &SimpleCache{
		lru: lru,
	}
	return c, nil
}

// Purge is used to completely clear the cache.
func (c *SimpleCache) Purge() {
	c.lock.Lock()
	c.lru.Purge()
	c.lock.Unlock()
}

// Add adds a value to the cache.  Returns true if an eviction occurred.
func (c *SimpleCache) Add(key, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.lru.Add(key, value)
}

// Get looks up a key's value from the cache.
func (c *SimpleCache) Get(key interface{}) (value interface{}, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Get(key)
}

// Contains checks if a key is in the cache, without updating the
// recent-ness or deleting it for being stale.
func (c *SimpleCache) Contains(key interface{}) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lru.Contains(key)
}

// Peek returns the key value (or undefined if not found) without updating
// the "recently used"-ness of the key.
func (c *SimpleCache) Peek(key interface{}) (value interface{}, ok bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lru.Peek(key)
}

// Probe adds 'val' if the key is NOT found in the cache and returns it.
// If key is in the cache, the corresponding value is returned.
// 'ok' is true is found in the cache and false otherwise.
func (c *SimpleCache) Probe(key interface{}, ctor func(key interface{}) interface{}) (value interface{}, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	value, ok = c.lru.Get(key)
	if !ok {
		// No hit; add to the cache
		value = ctor(key)

		c.lru.Add(key, value)
	}
	return value, ok
}

// ContainsOrAdd checks if a key is in the cache  without updating the
// recent-ness or deleting it for being stale,  and if not, adds the value.
// Returns whether found and whether an eviction occurred.
func (c *SimpleCache) ContainsOrAdd(key, value interface{}) (ok, evicted bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.lru.Contains(key) {
		return true, false
	}
	evicted = c.lru.Add(key, value)
	return false, evicted
}

// Remove removes the provided key from the cache.
func (c *SimpleCache) Remove(key interface{}) {
	c.lock.Lock()
	c.lru.Remove(key)
	c.lock.Unlock()
}

// RemoveOldest removes the oldest item from the cache.
func (c *SimpleCache) RemoveOldest() {
	c.lock.Lock()
	c.lru.RemoveOldest()
	c.lock.Unlock()
}

// Keys returns a slice of the keys in the cache, from oldest to newest.
func (c *SimpleCache) Keys() []interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lru.Keys()
}

// Len returns the number of items in the cache.
func (c *SimpleCache) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lru.Len()
}
