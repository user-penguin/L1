/*Package p8
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0
*/
package p8

import (
	"bufio"
	"fmt"
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

	// проверка, что в значение бита точно хотят установить 0 или 1
	if bit < 0 || bit > 1 {
		fmt.Printf("Значение бита может быть или 1, или 0, попробуйте ещё раз")
		return
	}

	fmt.Printf("Исходное число %d в прямом коде:\n %064b\n", sourceNum, uint64(sourceNum))
	changedSrc := changeBit(sourceNum, position, bit)
	fmt.Printf("Новое число %d в прямом коде:\n %064b\n", changedSrc, uint64(changedSrc))
}

func changeBit(src int64, position int64, value int64) int64 {
	tmp := uint64(src)
	mask := uint64(1 << position)
	switch value {
	case 0:
		return int64(tmp &^ mask)
	default:
		return int64(tmp | mask)
	}
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
