/*Package p13
Поменять местами два числа без создания временной переменной.
*/
package p13

import (
	"fmt"
)

func Run() {
	spin1()
	spin2()
}

func spin1() {
	a, b := 5, 13
	a, b = b, a
	printAB(a, b)
}

func spin2() {
	a, b := 6, 15
	a = a * b // a = 90
	b = a / b // b = 15
	a = a / b // a = 6
	printAB(a, b)
}

func printAB(a int, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}
