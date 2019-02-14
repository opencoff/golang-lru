# golang-lru

[![GoDoc](https://godoc.org/github.com/opencoff/golang-lru?status.svg)](https://godoc.org/github.com/opencoff/golang-lru)
[![Go Report Card](https://goreportcard.com/badge/github.com/opencoff/golang-lru)](https://goreportcard.com/report/github.com/opencoff/golang-lru)

This provides the `lru` package which implements a fixed-size
thread safe LRU cache. It is based on the cache in Groupcache.

## Changes from the original Hashicorp LRU Package
* Addition of a new `Probe()` ethod to each cache implementation.
  This enables caller to construct and add an element into the cache
  if it is not present.
* Renaming `lru.Cache` to `lru.SimpleCache`
* Addition of a new interface `lru.Cache` to abstract the
  implementation of 3 concrete cache types. Callers can use this
  interface to be insulated from specific implementation details.


# Example

Using the LRU is very simple:

```go

import "github.com/opencoff/golang-lru"


func foo() {
    var c lru.Cache
    var err error

    c, err = lru.NewSimple(128)
    if err != nil {
        panic(err)
    }

    for i := 0; i < 256; i++ {
        l.Add(i, nil)
    }

    if l.Len() != 128 {
        panic(fmt.Sprintf("bad len: %v", l.Len()))
    }
}
```
