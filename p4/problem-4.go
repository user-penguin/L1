/*Package p4
Задача 4.
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/
package p4

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//Структура со всеми каналами
type control struct {
	sig  chan os.Signal
	data chan string
	stop chan interface{}
	term chan interface{}
}

func NewControl(poolSize int) *control {
	return &control{
		sig:  make(chan os.Signal, 1),
		data: make(chan string, poolSize),
		stop: make(chan interface{}),
		term: make(chan interface{}),
	}
}

func (c *control) wait() {
	<-c.term
	close(c.term)
}

func Run() {
	workerPoolSize := 10
	c := NewControl(workerPoolSize)

	//подготавливаемся и запускаем горутину, которая будет ловить ^C
	signal.Notify(c.sig, syscall.SIGINT)
	go interruptControl(c)
	//запускаем запись и читающих воркеров
	go fillChannel(c)
	go startWorkers(workerPoolSize, c)

	c.wait()
}

func interruptControl(c *control) {
	for {
		switch <-c.sig {
		case syscall.SIGINT:
			fmt.Printf("Get ^C\n")
			c.stop <- 0
			close(c.stop)
		default:
			fmt.Printf("Unknown signal\n")
		}
	}
}

//Функция запускает работу воркеров в количестве workerPoolSize
//воркеры работают до закрытия канала с данными, пока не вычитают всё
func startWorkers(workerPoolSize int, c *control) {
	wg := new(sync.WaitGroup)
	for i := 0; i < workerPoolSize; i++ {
		workerNumber := i
		wg.Add(1)
		go work(c.data, workerNumber, wg)
	}
	wg.Wait()
	fmt.Println("Чтение завершено")
	c.term <- 0
}

//work запускает работу воркера до закрытия канала с данными
func work(dataCh chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataCh {
		fmt.Printf("worker(%d) data from channel: %s, len(%d) cap(%d)\n", id, data, len(dataCh), cap(dataCh))
		time.Sleep(time.Millisecond * 100)
	}
}

//fillChannel - функция, которая обеспечивает постоянное наполнение канала данными
//при получении значения из stopChan прекращает наполнять канал с данными и выходит из бесконечного фора
func fillChannel(c *control) {
	for {
		select {
		case <-c.stop:
			fmt.Println("Получен сигнал остановки")
			close(c.data)
			return
		case c.data <- generateData():
			time.Sleep(time.Millisecond)
		}
	}
}

//Генерация случайной строки случайной длины
func generateData() string {
	//случайная длина строки
	sizeMin, sizeMax := 2, 200
	size := rand.Intn(sizeMax-sizeMin) + sizeMin

	//массив байтов под результат
	res := make([]byte, size)
	for n := range res {
		res[n] = byte(rand.Intn('z'-'0') + '0')
	}
	return string(res)
}
