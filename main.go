package main

import (
	"L1/p8"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Какую задачу запускаем?")
	bitNumber, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("%s", err)
	}
	bitNumber = strings.TrimSuffix(bitNumber, "\n")
	taskNumber, err := strconv.ParseInt(bitNumber, 10, 64)
	if err != nil {
		fmt.Printf("Введённые символы не переводятся в int.\n Попробуйте ещё раз через пару минут.\n%s", err)
	}
	runTask(taskNumber)
}

func runTask(number int64) {
	switch number {
	case 8:
		p8.Run()
	default:
		fmt.Printf("Такая задача пока не оформлена\n")
	}
}
