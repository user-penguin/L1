/*Package p14
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/
package p14

import (
	"fmt"
	"reflect"
)

func Run() {
	var variable string
	n := 12
	var b bool
	ch := make(chan interface{})
	ch2 := make(chan int)
	fmt.Printf("The type of variable is: %s\n", printType(variable))
	fmt.Printf("The type of variable is: %s\n", printType(n))
	fmt.Printf("The type of variable is: %s\n", printType(b))
	fmt.Printf("The type of variable is: %s\n", printType(ch))
	fmt.Printf("The type of variable is: %s\n", printType(ch2))
}

func printType(variable interface{}) reflect.Type {
	return reflect.TypeOf(variable)
}
