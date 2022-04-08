/*Package p12
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
package p12

import "fmt"

func Run() {
	src := []string{
		"man",
		"cat",
		"man",
		"cat",
		"girl",
		"girl",
		"girl",
		"tree",
		"dog",
		"cat",
	}
	for _, str := range getOwnSet(src) {
		fmt.Printf("{%s} ", str)
	}
	println()
}

func getOwnSet(src []string) []string {
	res := make([]string, 0)
	checkMap := make(map[string]interface{})
	for _, str := range src {
		checkMap[str] = struct{}{}
	}
	for key := range checkMap {
		res = append(res, key)
	}
	return res
}
