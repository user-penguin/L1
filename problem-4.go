/*
Задача 4.
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	workerPoolSize := 10
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	go interruptControl(sigCh)
	run(workerPoolSize)
}

func interruptControl(sigCh chan os.Signal) {
	for {
		s := <-sigCh
		switch s {
		case syscall.SIGINT:
			fmt.Printf("Get ~C\n")
			os.Exit(0)
		default:
			fmt.Printf("Unknown signal\n")
		}

	}
}

func run(workerPoolSize int) {
	dataCh := make(chan string, workerPoolSize)
	go fillChannel(dataCh)
	wg := new(sync.WaitGroup)
	for i := 0; i < workerPoolSize; i++ {
		workerNumber := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range dataCh {
				fmt.Printf("worker(%d) data from channel: %s, len(%d) cap(%d)\n", workerNumber, data, len(dataCh), cap(dataCh))
				//time.Sleep(time.Millisecond * 1000)
			}
		}()
	}
	wg.Wait()
}

func fillChannel(dataChan chan string) {
	for i := 0; i < 1000; i++ {
		dataChan <- generateData()
		//time.Sleep(time.Millisecond * 1)
	}
	close(dataChan)
}

const alf = "qwertyuiopasdfghjklzxcvbnm1234567890"

func generateData() string {
	length := rand.Int63()%100 + 1 //случайная длина сообщения от 1 до 100
	bytes := make([]byte, length)  // массив байт под возвращаемое слово
	for i := range bytes {
		bytes[i] = alf[rand.Int63()%int64(len(alf))]
	}
	return string(bytes)
}
