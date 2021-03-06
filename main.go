package main

import (
	"L1/p1"
	"L1/p10"
	"L1/p11"
	"L1/p12"
	"L1/p13"
	"L1/p14"
	"L1/p15"
	"L1/p16"
	"L1/p17"
	"L1/p18"
	"L1/p19"
	"L1/p2"
	"L1/p20"
	"L1/p21"
	"L1/p22"
	"L1/p23"
	"L1/p24"
	"L1/p25"
	"L1/p26"
	"L1/p3"
	"L1/p4"
	"L1/p5"
	"L1/p6"
	"L1/p7"
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
	case 4:
		p4.Run()
	case 5:
		p5.Run()
	case 6:
		p6.Run()
	case 7:
		p7.Run()
	case 8:
		p8.Run()
	case 9:
		p9.Run()
	case 10:
		p10.Run()
	case 11:
		p11.Run()
	case 12:
		p12.Run()
	case 13:
		p13.Run()
	case 14:
		p14.Run()
	case 15:
		p15.Run()
	case 16:
		p16.Run()
	case 17:
		p17.Run()
	case 18:
		p18.Run()
	case 19:
		p19.Run()
	case 20:
		p20.Run()
	case 21:
		p21.Run()
	case 22:
		p22.Run()
	case 23:
		p23.Run()
	case 24:
		p24.Run()
	case 25:
		p25.Run()
	case 26:
		p26.Run()
	default:
		fmt.Printf("Такая задача пока не оформлена\n")
	}
}
