/*Package p9
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/
package p9

import (
	"fmt"
	"sync"
)

func Run() {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}

	chInput := make(chan int)
	chOutput := make(chan int)

	wg := new(sync.WaitGroup)
	wg.Add(len(nums))

	go writeSource(nums, chInput)
	go multi2(chInput, chOutput)
	go readOutput(chOutput, wg)

	wg.Wait()
}

// записываем из массива в первый канал
func writeSource(data []int, ch chan int) {
	defer close(ch)
	for _, elem := range data {
		ch <- elem
	}
}

//из второго канала пишем в консольку
func readOutput(ch chan int, wg *sync.WaitGroup) {
	for elem := range ch {
		fmt.Println(elem)
		wg.Done() // заканчивая с очередным числом, вычитаем 1 из счётчика
	}
}

// читаем из первого канала и пишем х2 во второй
func multi2(chIn chan int, chOut chan int) {
	defer close(chOut)
	for elem := range chIn {
		chOut <- elem * 2
	}
}
