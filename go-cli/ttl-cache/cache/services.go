package cache

import (
	"fmt"
	"time"
)


func (c *LRUCache) Set(key string, val any, ttl int64) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// calculate expiry time
	// expiresAt := time.Now().Unix() + ttl
	expiresAt := time.Now().Add(time.Duration(ttl) * time.Second).UnixNano()

	// check if key exists
	if element, ok := c.Items[key]; ok{
		c.EvictList.MoveToFront(element)

		// update data
		ent := element.Value.(*EntryCache)
		ent.value = val
		ent.expiresAt = expiresAt
		return
	}

	// if new key
	newEntry := &EntryCache{
		key: key,
		value: val,
		expiresAt: expiresAt,
	}
	newElement := c.EvictList.PushFront(newEntry)
	c.Items[key] = newElement

	// check if map ka size > capacity
	if len(c.Items) > c.Capacity{
		fmt.Println("Size limmit Exceeded!")
		lastNode := c.EvictList.Back()
		if lastNode != nil{
			// remove from list
			c.EvictList.Remove(lastNode)

			// remove from map 
			kv := lastNode.Value.(*EntryCache)
			delete(c.Items, kv.key)
		}
	}
}


func (c *LRUCache) Get(key string) (any, bool){
	c.mu.Lock()
	defer c.mu.Unlock()

	currentTime := time.Now().UnixNano()

	// element ka check karte hain exist karta hai k nhi
	if element, ok := c.Items[key]; ok{
		// TTL check happening
		enteredData := element.Value.(*EntryCache)

		// yaha se we check TTl expiry first
		// Lazy Expiry using for now
		if currentTime > enteredData.expiresAt{
			fmt.Println("Data Expired>>>>")
			c.EvictList.Remove(element)
			delete(c.Items, key)
			return nil, false
		}

		// LRU check 
		c.EvictList.MoveToFront(element)

		fmt.Println("CACHE HIT")
		fmt.Println("Data Found >>> ", element.Value.(*EntryCache).value)
		return element.Value, true
	}

	fmt.Println("CACHE MISS")
	fmt.Println("Data Does not Exist!!")
	return "-1",  false
}


func (c *LRUCache) Show(){
	for e:=c.EvictList.Front(); e!=nil; e=e.Next(){
		data := e.Value.(*EntryCache)
		fmt.Printf("[%v: %v] <-> ", data.key, data.value)
	}
	fmt.Println("nil")
}

