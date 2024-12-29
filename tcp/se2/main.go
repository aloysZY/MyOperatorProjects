package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s: %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)

	// 模拟处理延迟，等待1秒后再响应
	time.Sleep(10 * time.Second)

	// 响应客户端
	fmt.Fprintf(w, "Hello, client! Your request was received.")
}

func main() {
	http.HandleFunc("/", handleRequest) // 设置路由处理函数

	// 启动HTTP服务器
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting HTTP server: %v\n", err)
		return
	}
}
