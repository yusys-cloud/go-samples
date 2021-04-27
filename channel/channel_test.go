// Author: yangzq80@gmail.com
// Date: 2020-09-24
//
package channel

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestBase(t *testing.T) {
	c := make(chan int)
	defer close(c)
	go func() { c <- 3 + 4 }()
	i := <-c
	fmt.Println(i)
}

func TestLoop(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}

func worker(done chan bool) {
	time.Sleep(time.Second)
	// 通知任务已完成
	done <- true
	log.Println("worker done")
}
func TestMainWaiting(t *testing.T) {
	done := make(chan bool, 1)
	go worker(done)
	// 等待任务完成
	<-done
	log.Println("main end")
}
