/*Package p18
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
package p18

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

func Run() {
	workerPoolSize := 3
	ch := make(chan int, workerPoolSize)
	wgWrite := new(sync.WaitGroup)
	wgRead := new(sync.WaitGroup)
	wgRead.Add(1)
	counter := 0
	go counterHandler(&counter, ch, wgRead)
	for i := 0; i < workerPoolSize; i++ {
		wgWrite.Add(1)
		go proc("worker"+strconv.Itoa(i), ch, wgWrite)
	}
	wgWrite.Wait()
	close(ch)
	wgRead.Wait()
	fmt.Println("Общее количество итераций: ", counter)
}

func counterHandler(counter *int, ch chan int, wg *sync.WaitGroup) {
	for range ch {
		*counter++
	}
	wg.Done()
}

func proc(id string, ch chan int, wg *sync.WaitGroup) {
	iterations := rand.Uint64() % uint64(10000)
	for iterations > 0 {
		fmt.Printf("Работает %s. Осталось %d Итераций\n", id, iterations-1)
		ch <- 1
		iterations--
	}
	wg.Done()
}
