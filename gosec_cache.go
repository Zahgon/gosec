package gosec

import (
	"container/list"
	"sync"
)

var GlobalCache = NewLRUCache[any, any](1 << 16)

type LRUCache[K comparable, V any] struct {
	capacity  int
	items     map[K]*list.Element
	evictList *list.List
	lock      sync.Mutex
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *LRUCache[K, V]) Add(key K, value V) { _ = "STUB: not implemented"; return }

func (c *LRUCache[K, V]) removeOldest() { _ = "STUB: not implemented"; return }
