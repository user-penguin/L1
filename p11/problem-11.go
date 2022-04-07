/*Package p11
Реализовать пересечение двух неупорядоченных множеств
*/
package p11

import "fmt"

func Run() {
	numsA := []int{8, 12, 5, 7, 3, 6, 7, 15, 4} // 5 7 6 15 16 2 0 66
	numsB := []int{16, 8, 3, 4, 2, 12, 0, 4, 66}
	allMap := make(map[int]int)
	for _, elemA := range numsA {
		allMap[elemA] = 1
	}
	for _, elemB := range numsB {
		if _, ok := allMap[elemB]; ok {
			allMap[elemB] = allMap[elemB] + 1
		} else {
			allMap[elemB] = 1
		}
	}
	result := make([]int, 0)
	for key, value := range allMap {
		if value < 2 {
			result = append(result, key)
		}
	}
	fmt.Println(result)
}
