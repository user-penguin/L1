/*Package p6
Реализовать все возможные способы остановки выполнения горутины.
*/
package p6

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	stopChan := make(chan int)
	stopFlag := false
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go goroutine1(stopChan, wg)
	go goroutine2(&stopFlag, wg)

	time.Sleep(time.Second * 2)

	stopFlag = true
	stopChan <- 1
	close(stopChan)
	wg.Wait()
}

// первая горутина, остнавливаем каналом
func goroutine1(stopCh chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-stopCh:
			fmt.Println("Канальная становочка goroutine1")
			wg.Done()
			return
		default:
			fmt.Println("goroutine1, работаем")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

// вторая горутина, останавливаем флагом
func goroutine2(stopFlag *bool, wg *sync.WaitGroup) {
	for {
		if *stopFlag {
			fmt.Println("Остановочка флагом goroutine2")
			wg.Done()
			return
		} else {
			fmt.Println("goroutine2, работаем")
			time.Sleep(time.Millisecond * 100)
		}
	}
}
