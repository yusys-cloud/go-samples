// Author: yangzq80@gmail.com
// Date: 2020-09-01
// 已验证可代理中间件：mysql，http
//
package tcpproxy

import (
	"fmt"
	"io"
	"net"
	"testing"
)

var (
	dst_addr   string = "127.0.0.1:3306" //目标服务器地址
	local_addr string = "0.0.0.0:1000"   //本地监听地址
)

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
	go io.Copy(remote_tcp, tcpConn)
	go io.Copy(tcpConn, remote_tcp)
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
