/*Package p23
Удалить i-ый элемент из слайса
*/
package p23

import (
	"fmt"
)

func Run() {
	sliceSrc := make([]int, 20)
	for i := len(sliceSrc) - 1; i >= 0; i-- {
		sliceSrc[i] = i
	}
	fmt.Println(sliceSrc)
	fmt.Printf("len: %d, cap %d\n", len(sliceSrc), cap(sliceSrc))
	sliceSrc = remove(sliceSrc, 10)
	fmt.Println(sliceSrc)
	fmt.Printf("len: %d, cap %d\n", len(sliceSrc), cap(sliceSrc))
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]  // на место удаляемого пихаем последний
	return s[:len(s)-1] // возвращаем срез с 0 по предпоследний
}
