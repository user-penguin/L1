/*Package p25
Реализовать собственную функцию sleep
*/
package p25

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Printf("Спим почти по заводскому\n")
	coolSleep(2)
	fmt.Printf("Проснулись\n")
	fmt.Printf("Спим похуже\n")
	worstSleep(4)
	fmt.Printf("Проснулись\n")

}

func worstSleep(seconds int) {
	start := time.Now().Unix()
	for time.Now().Unix()-start < int64(seconds) {

	}
}

func coolSleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}
