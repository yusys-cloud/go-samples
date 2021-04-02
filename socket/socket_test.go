// Author: yangzq80@gmail.com
// Date: 2020-09-05
//
package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
)

func connHandler(c net.Conn) {
	if c == nil {
		return
	}
	buf := make([]byte, 4096)
	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		inputs := strings.Split(inStr, " ")
		switch inputs[0] {
		case "ping":
			c.Write([]byte("pong\n"))
		case "echo":
			echoStr := strings.Join(inputs[1:], " ") + "\n"
			fmt.Println(echoStr)
			c.Write([]byte(echoStr))
		case "quit":
			c.Close()
			break
		default:
			fmt.Printf("Unsupported command: %s\n", inputs[0])
		}
	}
	fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
}

func TestServer(t *testing.T) {
	server, err := net.Listen("tcp", ":1208")
	if err != nil {
		fmt.Printf("Fail to start server, %s\n", err)
	}
	fmt.Println("Server Started ...")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Fail to connect, %s\n", err)
			break
		}
		go connHandler(conn)
	}
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:1208")
	defer conn.Close()
	if err != nil {
		fmt.Printf("Fail to connect, %s\n", err)
		return
	}
	//clientConnHandler(conn)
	handleCmd(conn, "ping")
	handleCmd(conn, "echo abcc")
	//handleCmd(conn,"quit")
}
func handleCmd(c net.Conn, cmd string) {

	buf := make([]byte, 1024)
	c.Write([]byte(cmd))
	cnt, err := c.Read(buf)
	if err != nil {
		fmt.Printf("Fail to read data, %s\n", err)
	}
	fmt.Print(string(buf[0:cnt]))
}

func clientInputConnHandler(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "quit" {
			return
		}
		c.Write([]byte(input))
		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
			continue
		}
		fmt.Print(string(buf[0:cnt]))
	}
}
