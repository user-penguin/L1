package main

import (
	"fmt"
	"time"
)

func main() {
	N := 3
	// with range
	start := time.Now().Unix()
	ch := make(chan string)
	go func() {
		for data := range ch {
			fmt.Printf("%s\n", data)
		}
	}()
	for time.Now().Unix()-start < int64(N) {
		ch <- "12"
		time.Sleep(time.Millisecond * 100)
	}
}

//todo with mutex
