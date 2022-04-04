package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Сколько секунд будем работать?")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Ошибка чтения: %s", err)
	}
	text = strings.TrimSuffix(text, "\n")
	duration, err := strconv.Atoi(text)
	if err != nil {
		fmt.Printf("Введённые символы не переводятся в int.\n Попробуйте ещё раз через пару минут.\n%s", err)
		return
	}
	fmt.Printf("Вы указали: %dc.\n", duration)
	limit := time.Second * time.Duration(duration)
	// with range
	start := time.Now()
	ch := make(chan string)
	go func() {
		for data := range ch {
			fmt.Printf("%s\n", data)
		}
	}()
	for time.Since(start) < limit {
		ch <- "12"
		time.Sleep(time.Second)
	}
}

//todo with mutex
