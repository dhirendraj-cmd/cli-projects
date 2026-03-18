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


// management cache using lru patern(may change the name as it is more of a manager than just lruy)
type LRUCache struct{
	Capacity int
	Items map[any]*list.Element
	mu sync.RWMutex
	EvictList *list.List
}

// constructore for lru
func NewLRUCache(capacity int) *LRUCache{
	return &LRUCache{
		Capacity: capacity,
		Items: make(map[any]*list.Element),
		EvictList: list.New(),
	}
}