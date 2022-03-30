package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	sum := 0
	for _, number := range nums {
		wg.Add(1)
		sum = sum + sqCalc(number, wg)
	}
	wg.Wait()
	fmt.Printf("Sum of number's square = %d", sum)
}

func sqCalc(num int, wg *sync.WaitGroup) int {
	defer wg.Done()
	return num * num
}
