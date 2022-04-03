package main

import (
	"fmt"
	"time"
)

func main() {
	limit := time.Second * 2
	// with range
	start := time.Now()
	ch := make(chan string)
	go func() {
		for data := range ch {
			fmt.Printf("%s\n", data)
		}
	}()
	for time.Since(start) < limit {
		ch <- "12"
		time.Sleep(time.Second)
	}
}

//todo with mutex
