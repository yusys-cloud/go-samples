// Author: yangzq80@gmail.com
// Date: 2021-09-09
//
package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestWg(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		var inwg sync.WaitGroup
		inwg.Add(1)
		go func() {
			fmt.Println("Inner goroutine 1 done")
			inwg.Done()
		}()
		inwg.Wait()

		fmt.Println("goroutine 1 done")
		wg.Done()
	}()
	go func() {
		fmt.Println("goroutine 2 done")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("all work is done")
}

func TestWaitGroupLoop(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {

		go func() {
			wg.Add(1)
			fmt.Println("goroutine done", i)
			wg.Done()
		}()

		time.Sleep(time.Second)
	}
	wg.Wait()
	fmt.Println("all work is done")
}

func TestTmp(t *testing.T) {
	for {
		fmt.Println(rand.Intn(3))
	}

	fmt.Println(11 / 5)
}
