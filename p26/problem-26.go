/*Package p26
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/
package p26

import (
	"fmt"
	"strings"
)

func Run() {
	src := []string{
		"abababa",
		"aAb",
		"abcd",
		"abCdefAaf",
		"aabcd",
	}
	for _, str := range src {
		fmt.Printf("Строка %s уникальна? - %v\n", str, isUniqLetters(str))
	}
}

func isUniqLetters(str string) bool {
	runes := []rune(strings.ToLower(str))
	check := make(map[rune]bool)
	for _, lett := range runes {
		if _, ok := check[lett]; ok {
			return false
		} else {
			check[lett] = true
		}
	}
	return true
}
