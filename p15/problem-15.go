/*Package p15
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/
package p15

import "fmt"

var justString string

func someFunc() {
	v := createHugeString(1 << 10) //1024
	justString = v[:100]
}

func createHugeString(size int) string {
	fmt.Println(size)
	res := ""
	for i := 0; i < size; i++ {
		res += "a"
	}
	return res
}

func Run() {
	someFunc()
	println(justString)
}
