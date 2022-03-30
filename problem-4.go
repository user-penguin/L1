package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	workerPoolSize := 10
	run(workerPoolSize)
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
				time.Sleep(time.Millisecond * 1000)
			}
		}()
	}
	wg.Wait()
	close(dataCh)
}

func fillChannel(dataChan chan string) {
	for {
		dataChan <- generateData()
		time.Sleep(time.Millisecond * 1)
	}
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
