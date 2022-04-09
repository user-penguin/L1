/*Package p22
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/
package p22

import (
	"fmt"
	"math/big"
)

func Run() {
	a, _ := new(big.Int).SetString("123000000000000000000000000000000", 10)
	b, _ := new(big.Int).SetString("234000000000000444400000000000000", 10)
	fmt.Printf("%v + %v = %v\n", a, b, Add(a, b))
	fmt.Printf("%v - %v = %v\n", a, b, Sub(a, b))
	fmt.Printf("%v * %v = %v\n", a, b, Mul(a, b))
	fmt.Printf("%v / %v = %v\n", a, b, Div(a, b))
}

func Add(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Sub(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func Mul(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Div(a *big.Int, b *big.Int) *big.Float {
	f1 := new(big.Float).SetInt(a)
	f2 := new(big.Float).SetInt(b)
	return new(big.Float).Quo(f1, f2)
}
