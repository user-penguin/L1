package main

import (
	"L1/p1"
	"L1/p10"
	"L1/p13"
	"L1/p2"
	"L1/p3"
	"L1/p5"
	"L1/p6"
	"L1/p8"
	"L1/p9"
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
	case 1:
		p1.Run()
	case 2:
		p2.Run()
	case 3:
		p3.Run()
	case 5:
		p5.Run()
	case 6:
		p6.Run()
	case 8:
		p8.Run()
	case 9:
		p9.Run()
	case 10:
		p10.Run()
	case 13:
		p13.Run()
	default:
		fmt.Printf("Такая задача пока не оформлена\n")
	}
}
