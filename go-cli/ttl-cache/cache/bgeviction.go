package cache

import (
	"container/heap"
	"fmt"
	"time"
)


func (c *LRUCache) BackgroundEviction(stopChan chan struct{}){
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

		for {
			select {
			case <- ticker.C:
				c.mu.Lock()
				c.mu.Unlock()
				currentTime := time.Now().UnixNano()

				for c.Expiration.Len() > 0{
					// peek at root which is at index 0
					root := (*c.Expiration)[0]

					if currentTime < root.expiresAt{
						break
					}

					// if currenttime > expires time
					expiredItem := heap.Pop(c.Expiration).(*ExpiryItem)

					// need to remove from map and list both
					if ele, ok := c.Items[expiredItem.key]; ok{
						// remove from list(lru)
						c.EvictList.Remove(ele)

						// delete from dictionary also
						delete(c.Items, expiredItem.key)
						fmt.Println("Background: Cleaned up >> ", expiredItem.key)
					}
				}
			
			case <- stopChan:
				fmt.Println("Stopping the background cleanup process gracefully!")
				return
			}
			
		}
}

