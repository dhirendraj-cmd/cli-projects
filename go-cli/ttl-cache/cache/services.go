package cache

import (
	"container/heap"
	"fmt"
	"time"
)


func (c *CacheManager) Set(key string, val any, ttl int64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// calculate expiry time
	// expiresAt := time.Now().Unix() + ttl
	expiresAt := time.Now().Add(time.Duration(ttl) * time.Second).UnixNano()

	// check if key exists
	if element, ok := c.Items[key]; ok{

		fmt.Println("Updating the entry for: ", key)
		// update data
		ent := element.Value.(*EntryCache)
		ent.value = val
		ent.expiresAt = expiresAt

		// update heap if key exists
		c.Expiration.Update(ent.expiryRef, ent.expiresAt)
		
		// lru call to move to front
		c.EvictList.MoveToFront(element)
		return
	}

	// if new key

	// first insert into ExpiryItem struct
	newExpiryItem := &ExpiryItem{
		key: key,
		expiresAt: expiresAt,
	}

	// push to heap library
	heap.Push(c.Expiration, newExpiryItem)

	/* Note to self: for pushing to heap no need to call the Programs Push function in cache.go as the main library heap.Push() calls it internally and then append the item at the end and call the Up() bcoz our Push() will just append at the end of slice but won't heapify it*/

	// new entry in cache 
	newEntry := &EntryCache{
		key: key,
		value: val,
		expiresAt: expiresAt,
		expiryRef: newExpiryItem,
	}

	// push to LRU
	newElement := c.EvictList.PushFront(newEntry)
	c.Items[key] = newElement

	// check if map ka size > capacity
	if len(c.Items) > c.Capacity{
		fmt.Println("Size limmit Exceeded!")
		lastNode := c.EvictList.Back()
		if lastNode != nil{
			// remove from heap
			kv := lastNode.Value.(*EntryCache)

			if kv.expiryRef != nil{
				heap.Remove(c.Expiration, kv.expiryRef.index)
			}
			// remove from list
			c.EvictList.Remove(lastNode)

			// remove from map 
			delete(c.Items, kv.key)
		}
	}
}


func (c *CacheManager) Get(key string) (any, bool){
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

			// remove from heap
			if enteredData.expiryRef != nil{
				heap.Remove(c.Expiration, enteredData.expiryRef.index)
			}

			// remove from lru
			c.EvictList.Remove(element)

			// remove from hashmap
			delete(c.Items, key)
			return nil, false
		}

		// LRU check 
		c.EvictList.MoveToFront(element)

		fmt.Println("CACHE HIT")
		fmt.Println("Data Found >>> ", element.Value.(*EntryCache).value)
		return enteredData.value, true
	}

	fmt.Println("CACHE MISS")
	fmt.Println("Data Does not Exist!!")
	return "-1",  false
}


func (c *CacheManager) Show(){
	for e:=c.EvictList.Front(); e!=nil; e=e.Next(){
		data := e.Value.(*EntryCache)
		conNow := time.Now().UnixNano()
		remNano := data.expiresAt - conNow
		remSecs := remNano / int64(time.Second)
		fmt.Printf("[%v: %v: %vs] <-> ", data.key, data.value, remSecs)
	}
	fmt.Println("nil")
}

