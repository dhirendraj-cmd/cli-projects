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
	Items map[any]*list.Element
	mu sync.RWMutex
	EvictList *list.List
}

func NewLRUCache(capacity int) *LRUCache{
	return &LRUCache{
		Capacity: capacity,
		Items: make(map[any]*list.Element),
		EvictList: list.New(),
	}
}