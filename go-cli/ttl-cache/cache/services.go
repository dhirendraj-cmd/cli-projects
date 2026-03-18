package cache

import (
	"fmt"
	"time"
)


func (c *LRUCache) Set(key any, val any, ttl int64) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// calculate expiry time
	expiresAt := time.Now().Unix() + ttl

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

	// check if size > capacity
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


