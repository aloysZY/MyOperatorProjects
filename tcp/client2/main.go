package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 拨号到服务器
	conn, err := net.Dial("tcp", "10.211.55.21:9090") // 请根据实际情况修改地址和端口
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error dialing: %s\n", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// 确保 conn 是 *net.TCPConn 类型
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		fmt.Fprintf(os.Stderr, "Connection is not a TCP connection.\n")
		os.Exit(1)
	}
	// 使用goroutine发送数据
	go func() {
		message := "GET / HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n" // 使用HTTP请求作为例子
		_, err := fmt.Fprintf(tcpConn, message)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to server: %s\n", err.Error())
		}
	}()

	// 先关闭写入方向
	err = tcpConn.CloseWrite()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending FIN: %s\n", err.Error())
	}

	// 立即关闭读取方向
	tcpConn.CloseRead()

	fmt.Println("Client finished.")
}
