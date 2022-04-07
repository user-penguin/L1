/*Package p3
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/
package p3

import (
	"fmt"
	"sync"
)

func Run() {
	nums := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	sum := 0
	for _, number := range nums {
		wg.Add(1)
		number := number
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			sum = sum + sqCalc(number)
		}()
	}
	wg.Wait()
	fmt.Printf("Sum of number's square = %d", sum)
}

func sqCalc(num int) int {
	return num * num
}
