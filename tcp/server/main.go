package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from client: %v\n", err)
			break
		}
		fmt.Printf("Received from client: %s", message)

		// 将接收到的消息立即发送回给客户端
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error writing to client: %v\n", err)
			break
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":18080")
	if err != nil {
		fmt.Printf("Error starting TCP server: %v\n", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 18080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go handleConnection(conn) // 启动一个goroutine来处理每个新的连接
	}
}
