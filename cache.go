package lru

// Cache is an interface describing a generic LRU cache. It has 3 concrete
// implementations:
//   * Simple LRU
//   * 2 Queue
//   * ARC (Adaptive Replacement Cache)
type Cache interface {
	// Add adds a value to the cache.
	Add(key, val interface{})

	// Get looks up a key's value from the cache.
	Get(key interface{}) (value interface{}, ok bool)

	// Probe adds 'val' if the key is NOT found in the cache and returns it.
	// If key is in the cache, the corresponding value is returned.
	// 'ok' is true is found in the cache and false otherwise.
	Probe(key interface{}, ctor func(key interface{}) interface{}) (value interface{}, ok bool)

	// Peek is used to inspect the cache value of a key
	// without updating recency or frequency.
	Peek(key interface{}) (value interface{}, ok bool)

	// Remove removes the provided key from the cache.
	Remove(key interface{})

	// Contains is used to check if the cache contains a key
	// without updating recency or frequency.
	Contains(key interface{}) bool

	// Len returns the number of items in the cache.
	Len() int

	// Keys returns a slice of the keys in the cache.
	// The frequently used keys are first in the returned slice.
	Keys() []interface{}

	// Purge is used to completely clear the cache.
	Purge()
}

// make sure each of these implement the interface methods
var _ Cache = &SimpleCache{}
var _ Cache = &TwoQueueCache{}
var _ Cache = &ARCCache{}
