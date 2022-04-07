/*Package p20
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
package p20

import (
	"fmt"
	"regexp"
	"strings"
)

func Run() {
	strSrc := "snow dog  sun,hhj"
	println(reverse(strSrc))
}

func reverse(str string) string {
	exp := regexp.MustCompile(`[\s,-]+`)
	splitStrs := exp.Split(str, -1)
	fmt.Println(splitStrs)
	strRes := ""
	for i := len(splitStrs) - 1; i >= 0; i-- {
		fmt.Println(strRes)
		strRes += splitStrs[i] + " "
	}
	return strings.TrimSuffix(strRes, " ")
}
