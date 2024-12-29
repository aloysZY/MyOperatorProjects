package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	// 创建自定义的HTTP客户端
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			// 设置写入超时时间为1毫秒，以确保在发送数据时遇到阻塞时迅速断开连接
			WriteBufferSize: 32 * 1024,
		},
	}

	// 创建HTTP请求
	req, err := http.NewRequest("GET", "http://10.211.55.21:9090/", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating request: %s\n", err.Error())
		os.Exit(1)
	}

	// 设置请求头
	req.Header.Set("Host", "example.com")
	req.Header.Set("Connection", "close")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending request: %s\n", err.Error())
	} else {
		// 立即关闭响应体，模拟客户端取消请求
		resp.Body.Close()
	}

	fmt.Println("Client finished.")
}
