/*Package p13
Поменять местами два числа без создания временной переменной.
*/
package p13

import (
	"fmt"
)

func Run() {
	a := 5
	b := 13
	a, b = a*b/a, a*b/b
	fmt.Printf("a = %d, b = %d\n", a, b)
}
