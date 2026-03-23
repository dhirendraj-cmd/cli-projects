package cache

import (
	"container/heap"
	"container/list"
	"sync"
)


type EntryCache struct{
	key string
	value any
	expiresAt int64
	expiryRef *ExpiryItem
}

type ExpiryItem struct {
	key       string
	expiresAt int64
	index     int
}

// heapify in go
type ExpiryHeap []*ExpiryItem

func (eh ExpiryHeap) Len() int { return len(eh)}

// inbuilt min heap logic calling for expiry check
func (eh ExpiryHeap) Less(i, j int) bool{
	return eh[i].expiresAt < eh[j].expiresAt
}

func (eh ExpiryHeap) Swap(i, j int) {
	eh[i], eh[j] = eh[j], eh[i]
	eh[i].index = i
	eh[j].index = j
}


// push to heap
func (eh *ExpiryHeap) Push(x any){
	n := len(*eh)
	item := x.(*ExpiryItem)
	item.index = n
	*eh = append(*eh, item)
}


// pop expired
func (eh *ExpiryHeap) Pop() any{
	old := *eh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*eh = old[0:n-1]
	return item
}

// fix min heap after update
func (eh *ExpiryHeap) Update(item *ExpiryItem, expiresAt int64){
	item.expiresAt = expiresAt
	heap.Fix(eh, item.index)
}

// management cache using lru patern(may change the name as it is more of a manager than just lruy)
type LRUCache struct{
	Capacity int
	Items map[string]*list.Element
	mu sync.RWMutex
	EvictList *list.List
	Expiration *ExpiryHeap
	HeapItems map[string]*ExpiryItem
}

// constructore for lru
func NewLRUCache(capacity int, stopChan chan struct{}) *LRUCache{
	h := &ExpiryHeap{}
	heap.Init(h)
	c := &LRUCache{
		Capacity: capacity,
		Items: make(map[string]*list.Element),
		EvictList: list.New(),
		Expiration: h,
		HeapItems: make(map[string]*ExpiryItem),
	}

	go c.BackgroundEviction(stopChan)
	return c
}

