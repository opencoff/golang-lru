package lru

// Cache is an interface describing a generic LRU cache. It has 3 concrete
// implementations:
//   * Simple LRU
//   * 2 Queue
//   * ARC (Adaptive Replacement Cache)
type Cache interface {
	// Add adds a value to the cache.
	Add(key, val interface{})

	// Get returns (value, true) if key is in the cache and
	// (nil, false) otherwise. If key is found in the cache, depending
	// on the implementation, the cache may update the frequency or recency
	// of the key.
	Get(key interface{}) (value interface{}, ok bool)

	// Probe returns (value, true) if key is already in the cache.
	// Otherwise, it constructs the new value by calling ctor(key), and
	// inserts the new value into the cache; finally it returns the
	// newly constructed value and false (value, false).
	Probe(key interface{}, ctor func(key interface{}) interface{}) (value interface{}, ok bool)

	// Peek is similar to Get() except, it doesn't update the
	// recency or frequency.
	Peek(key interface{}) (value interface{}, ok bool)

	// Remove removes key from the cache if it exists.
	Remove(key interface{})

	// Contains returns true if the key is found in the cache and
	// false otherwise. It does this operation without updating recency
	// or frequency.
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
