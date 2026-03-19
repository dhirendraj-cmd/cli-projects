package main

import (
	"fmt"
	"github.com/dhirendraj-cmd/cli-projects/tree/main/go-cli/ttl-Cache/cache"
)

func main(){
	fmt.Println("CLI Based TTL Cache")
	// cache.MiniTTLLRUCache()
	lru := cache.NewLRUCache(5)
	lru.Set(1, "AB", 50)
	lru.Set(2, "BC", 50)
	lru.Set(3, "CD", 50)
	lru.Set(4, "DE", 30)

	fmt.Println("GET DATA >>>>>>> ")
	lru.Get(3)
}

