/*Package p7
Реализовать конкурентную запись данных в map.
*/
package p7

import (
	"fmt"
	"strconv"
	"sync"
)

type safeMap struct {
	mx sync.Mutex
	m  map[string]int
}

func (s *safeMap) addValue(key string, value int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.m[key] = value
}

func (s *safeMap) readValue(key string) (int, bool) {
	s.mx.Lock()
	defer s.mx.Unlock()
	val, ok := s.m[key]
	return val, ok
}

func NewSafeMap() *safeMap {
	return &safeMap{
		m: make(map[string]int),
	}
}

func Run() {
	store := NewSafeMap()
	wg := new(sync.WaitGroup)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go writeSquare(i, store, wg)
	}
	wg.Done()
	fmt.Println(store)
}

func writeSquare(i int, store *safeMap, wg *sync.WaitGroup) {
	sq := calcSquare(i)
	store.addValue(strconv.Itoa(i), sq)
	wg.Done()
}

func calcSquare(i int) int {
	return i * i
}
