// Author: yangzq80@gmail.com
// Date: 2020-09-01
// 已验证可代理中间件：mysql，http
//
package tcpproxy

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	dst_addr   string = "127.0.0.1:3306" //目标服务器地址
	local_addr string = "0.0.0.0:1000"   //本地监听地址
)

// 测试通过1000端口访问mysql数据库
func TestTCPProxy(t *testing.T) {
	fmt.Println("this is porxy server:")
	go tcp_listen()

	select {}
}

func tcp_handle(tcpConn net.Conn) {
	remote_tcp, err := net.Dial("tcp", dst_addr) //连接目标服务器

	if err != nil {
		fmt.Println(err)
		return
	}
	//go io.Copy(remote_tcp, tcpConn)
	go CopyBuffer(remote_tcp, tcpConn, nil)
	//go io.Copy(tcpConn, remote_tcp)
	go CopyBuffer(tcpConn, remote_tcp, nil)
}

func tcp_listen() {
	ln, err := net.Listen("tcp", local_addr)
	if err != nil {
		fmt.Println("tcp_listen:", err)
		return
	}
	defer ln.Close()
	for {
		tcp_Conn, err := ln.Accept() //接受tcp客户端连接，并返回新的套接字进行通信
		if err != nil {
			fmt.Println("Accept:", err)
			return
		}
		go tcp_handle(tcp_Conn) //创建新的协程进行转发
	}
}

//测试socket 读写内容
func TestSocket(t *testing.T) {
	go startServer()

	if len(os.Args) <= 1 {
		fmt.Println("usage: go run client2.go YOUR_CONTENT")
		return
	}
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")

	time.Sleep(time.Second * 2)
	//data := os.Args[1]
	data := 1
	for {
		data++

		time.Sleep(time.Second * 2)

		conn.Write([]byte("aaaasdds1111222" + strconv.Itoa(data)))
	}

	time.Sleep(time.Second * 10000)
}

func startServer() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		var buf = make([]byte, 10)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
