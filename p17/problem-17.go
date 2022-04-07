/*Package p17
Реализовать бинарный поиск встроенными методами языка
*/
package p17

import "fmt"

func Run() {
	source := []int{-15, -2, 1, 2, 5, 6, 23, 33, 34}
	find := 2
	index := binarySearch(source, find)
	if index >= 0 {
		fmt.Printf("Элемент %d найден, его позиция в массиве %d\n", find, index)
	} else {
		fmt.Printf("Элемент %d не найден", find)
	}
}

// бинарный поиск работает для упорядоченного массива!
func binarySearch(src []int, find int) int {
	first, last := 0, len(src)-1
	for first <= last {
		mid := (first + last) / 2
		switch {
		case src[mid] == find:
			return mid
		case src[mid] < find:
			first = mid + 1
		case src[mid] > find:
			last = mid - 1
		}
	}
	return -1
}
