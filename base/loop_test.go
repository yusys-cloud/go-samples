// Author: yangzq80@gmail.com
// Date: 2021-09-03
//
package base

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for {
		select {
		case <-ticker.C:
			i++
			log.Println(i)
			if i == 10 {
				ticker.Stop()
				return
			}
		}

	}
}

func TestTimer(t *testing.T) {

	stopper := time.NewTimer(time.Second * 4)

	running := true

	go func() {
		<-stopper.C
		running = false
		fmt.Println("4s after")
	}()

	ticker := time.NewTicker(500 * time.Millisecond)
	for running {
		<-ticker.C
		log.Println("11")
	}
}
