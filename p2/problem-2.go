/*Package p2
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package p2

import (
	"fmt"
	"sync"
)

func Run() {
	numbers := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	for _, num := range numbers {
		wg.Add(1)
		go printSquare(num, wg)
	}
	wg.Wait()
}

func printSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d * %d = %d\n", num, num, calcSquare(num))
}

func calcSquare(num int) int {
	return num * num
}
