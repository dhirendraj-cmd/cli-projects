package main

import (
	"cmd/cmd/uploads"
	"fmt"
)

func main(){
	fmt.Println("Goroutines projects")
	uploads.MoveFiles()
}
