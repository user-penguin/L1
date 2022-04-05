/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go goroutine1(ch)
	time.Sleep(time.Second * 2)
	ch <- 1
	close(ch)
}

// первая горутина, остнавливаем каналом
func goroutine1(ch chan int) {
	for {
		select {
		case <-ch:
			fmt.Println("Остановочка")
			return
		default:
			fmt.Println("Метро люблино, работаем")
			time.Sleep(time.Millisecond * 100)
		}
	}
}
