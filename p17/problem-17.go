/*Package p17
Реализовать бинарный поиск встроенными методами языка
*/
package p17

import "fmt"

func Run() {
	source := []int{6, 1, 23, 33, -2, 5, 34, -15, 6, 2}
	find := 2
	index := binarySearch(source, find)
	if index >= 0 {
		fmt.Printf("Элемент %d найден, его позиция в массиве %d\n", find, index)
	} else {
		fmt.Printf("Элемент %d не найден", find)
	}
}

func binarySearch(src []int, find int) int {
	return -1
}
