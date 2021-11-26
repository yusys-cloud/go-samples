// Author: yangzq80@gmail.com
// Date: 2020-09-24
// 通道（channel）是用来传递数据的一个数据结构
// 用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯
// 操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道
// ch := make(chan int) 不带缓存区
// ch := make(chan int,2) 指定缓冲区大小
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

func TestSize(t *testing.T) {
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	//如果没有缓存区则不能发送数据
	//go func() {
	//	ch <- 1
	//	ch <- 2
	//}()

	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func TestNoBuffer(t *testing.T) {
	ch := make(chan int)

	//如果没有缓存区则不能同步发送数据
	go func() {
		ch <- 1
		ch <- 2
	}()

	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//============struct channel=========
type Data struct {
	i int
}

func func1(c chan *Data) {
	for {
		var t *Data
		t = <-c   //receive
		t.i += 10 //increment
		c <- t    //send it back
	}
}

func TestChannelStruct(t *testing.T) {
	c := make(chan *Data)
	d := Data{10}
	go func1(c)
	println(d.i)
	c <- &d  //send a pointer to our t
	i := <-c //receive the result
	time.Sleep(1 * time.Second)
	println(i.i)
	println(d.i, "==")
}

//====ticker===
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Println(ticker.C)
		}
	}

}
