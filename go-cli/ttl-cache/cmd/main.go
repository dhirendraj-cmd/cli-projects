package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
	"github.com/dhirendraj-cmd/cli-projects/tree/main/go-cli/ttl-Cache/cache"
)

func main(){
	fmt.Println("CLI Based TTL Cache")
	// cache.MiniTTLLRUCache()

	stopChan := make(chan struct{})
	miniCache := cache.NewCacheManager(5, stopChan)

	// wait till program exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("TTL Cache CLI Started (Type 'exit' to quit)")
	fmt.Println("Commands: set <key> <val> <ttl_seconds>, get <key>, delete <key>, show")

	go func ()  {
		// wait till ctrl+c not prerssed
		<-sigChan
		fmt.Println("Shutting Gracefully!!")
		close(stopChan)
		time.Sleep(500 * time.Millisecond)
		os.Exit(0)
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() { break }

		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) == 0 { 
			continue
		}

		command := strings.ToLower(args[0])

		switch command {
		case "exit":
			return
		case "set":
			if len(args) < 4{
				fmt.Println("Usage: set <key> <value> <ttl_seconds>")
				continue
			}
			key := args[1]
			val := args[2]

			// parsing to Int for ttl
			ttl, err := strconv.ParseInt(args[3], 10, 64)

			if err != nil {
				fmt.Println("Invalid TTL. Must be a number.")
				continue
			}

			miniCache.Set(key, val, ttl)
			fmt.Printf("OK: %s stored with TTL %ds\n", key, ttl)
		case "get":
			if len(args) < 2{
				fmt.Println("Usage: get <key>")
				continue
			}
			val, found := miniCache.Get(args[1])
			if found{
				fmt.Printf("Value: %v\n", val)
			} else {
				fmt.Println("Data Not Found or Expired")
			}

		case "show":
			miniCache.Show()

		default:
			fmt.Println("?? Unknown command. Try: set, get, delete, show, exit")
		}
	}

}

// val := strings.Join(args[2:len(args)-1], " ")