/*Package p19
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/
package p19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку?")
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("%s", err)
	}
	str = strings.TrimSuffix(str, "\n")
	fmt.Println(reverseStr(str))
}

func reverseStr(str string) string {
	runeStr := []rune(str)
	resByte := make([]rune, len(runeStr))
	for i := len(runeStr) - 1; i >= 0; i-- {
		resByte = append(resByte, runeStr[i])
	}
	return string(resByte)
}
