/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0
*/
package p8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Run() {
	sourceNum, err := readInt("Над каким числом экспериментируем?")
	if err != nil {
		fmt.Printf("Ошибка во время чтения: %s", err)
	}
	position, err := readInt("Какой бит будем менять? (Давайте считать с нулевого)")
	if err != nil {
		fmt.Printf("Ошибка во время чтения: %s", err)
	}
	bit, err := readInt("Что запишем в бит?")
	if err != nil {
		fmt.Printf("Ошибка во время чтения: %s", err)
	}
	if bit < 0 || bit > 1 {
		fmt.Printf("Значение бита может быть или 1, или 0, попробуйте ещё раз")
		return
	}

	fmt.Printf("%d in binary = %64b\n", sourceNum, uint64(sourceNum))

	if bit == 1 {
		// надо получить 000010
		mask := pow2(position)
		res := sourceNum | mask
		fmt.Printf("Результат %064b %d\n", res, res)
	} else {
		// надо получить 111101
		mask := math.MaxInt64 &^ pow2(position)
		res := sourceNum & mask
		fmt.Printf("Результат %064b %d\n", res, res)
	}
}

func pow2(value int64) int64 {
	var res, i int64 = 1, 0
	for i = 0; i < value; i++ {
		res *= 2
	}
	return res
}

func readInt(msg string) (int64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(msg)
	bitNumber, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	bitNumber = strings.TrimSuffix(bitNumber, "\n")
	result, err := strconv.ParseInt(bitNumber, 10, 64)
	if err != nil {
		fmt.Printf("Введённые символы не переводятся в int.\n Попробуйте ещё раз через пару минут.\n%s", err)
		return 0, err
	}
	return result, nil
}
