/*Package p5
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
package p5

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Run() {
	// читаем продолжительность работы программы
	duration, err := stdInputReadInt("Сколько секунд будем работать?")
	if err != nil {
		panic(err)
	}
	msgPool := []string{
		"Добрый день",
		"Как дела",
		"Я есть горутина",
		"Да кто такие эти Роннины",
		"Семь раз отмерь, ну и хватит",
		"Это снова я, привет",
		"Всё очень, очень плохо",
	}
	// канал для сообщений
	msgCh := make(chan string)
	defer close(msgCh)

	// устанавливаем лимит времени в секундах
	limit := time.Second * time.Duration(duration)
	// засекаем время
	start := time.Now()

	go readCh(msgCh)

	// последовательно отправляем данные в канал
	for time.Since(start) < limit {
		msgCh <- msgPool[rand.Uint64()%uint64(len(msgPool))]
		time.Sleep(time.Second)
	}
}

func readCh(msg chan string) {
	for text := range msg {
		fmt.Printf("Получено сообщение: %s\n", text)
	}
}

func stdInputReadInt(msg string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(msg)
	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	text = strings.TrimSuffix(text, "\n")
	duration, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	return duration, nil
}
