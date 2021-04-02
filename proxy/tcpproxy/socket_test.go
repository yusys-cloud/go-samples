// Author: yangzq80@gmail.com
// Date: 2020-09-08
// net.TCPConn是允许在TCP客户端和TCP服务器之间的全双工通信的Go类型。两种主要方法是
// func (c *TCPConn) Write(b []byte) (n int, err os.Error)
// func (c *TCPConn) Read(b []byte) (n int, err os.Error)
// TCPConn被客户端和服务器用来读写消息。
//
package tcpproxy

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "3333", "port")

func TestClientCallServer(t *testing.T) {
	go startServer1()

	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + *port)

	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()

	time.Sleep(time.Second * 2)
}

func startServer1() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + *host + ":" + *port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest1(conn)
	}
}

// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
func CopyBuffer(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	//if wt, ok := src.(io.WriterTo); ok {
	//	return wt.WriteTo(dst)
	//}
	//// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	//if rt, ok := dst.(io.ReaderFrom); ok {
	//	return rt.ReadFrom(src)
	//}
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*io.LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	i := 1
	for {
		i++
		nr, er := src.Read(buf)
		log.Println(i, "=====================", string(buf))
		//time.Sleep(time.Millisecond*100)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}

func handleRequest1(conn net.Conn) {
	defer conn.Close()
	//io.Copy(conn, conn)
	CopyBuffer(conn, conn, nil)
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 10; i > 0; i-- {
		_, e := conn.Write([]byte("hello " + strconv.Itoa(i) + "\r\n"))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
		//fmt.Println("send...",i)
	}
}
func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}
	//fmt.Println("handle-read",*(*string)(unsafe.Pointer(&buf)),string(buf[:reqLen-1]))
	fmt.Println("handle-read", string(buf[:reqLen]))
}
