package cache

import (
	"container/heap"
	"fmt"
	"time"
)


func (c *CacheManager) BackgroundEviction(stopChan chan struct{}){
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <- ticker.C:
			c.evictExpiredItems() // seperate scoped function to handle locking correctly

		case <- stopChan:
			fmt.Println("Stopping the background cleanup process gracefully!")
			return
		}
		
	}
}


func(c *CacheManager) evictExpiredItems(){
	c.mu.Lock()
	defer c.mu.Unlock()

	currentTime := time.Now().UnixNano()
	evictCount := 0
	maxEvictionsPerBatch := 100 // prevent loading for too long so blocking doesn't happen

	for c.Expiration.Len() > 0{
		// limit batch size to keep the cache responsive
		if evictCount > maxEvictionsPerBatch { break }

		// peak at root which is at index 0
		root := (*c.Expiration)[0]

		if currentTime < root.expiresAt { break }

		// expired < currenttime
		expiredItem := heap.Pop(c.Expiration).(*ExpiryItem)

		// remove from map and DLL both
		if ele, ok := c.Items[expiredItem.key]; ok{

			// remove from list
			c.EvictList.Remove(ele)

			// remove from map
			delete(c.Items, expiredItem.key)
			fmt.Println("Background: Cleaned up >> ", expiredItem.key)
		}
	}

}

