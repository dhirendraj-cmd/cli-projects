package cache

import (
	"container/list"
	"sync"
)


type EntryCache struct{
	key any
	value any
	expiresAt int64
}

type LRUCache struct{
	Capacity int
	Items map[string]*list.Element
	mu sync.RWMutex
	EvictList *list.List
}

func NewLRUCache(capacity int) *LRUCache{
	return &LRUCache{
		Capacity: capacity,
		Items: make(map[string]*list.Element),
		EvictList: list.New(),
	}
}