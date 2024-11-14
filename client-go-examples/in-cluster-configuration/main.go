package main

import (
	"context"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// 获取集群内配置
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}
	// 根据配置信息创建client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	// 循环读取并打印Pod列表
	for {
		pods, err := clientSet.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("There are %d pods in the cluster\n", len(pods.Items))
		for i, pod := range pods.Items {
			log.Printf("%d -> %s/%s\n", i+1, pod.Namespace, pod.Name)
		}
		// 定时5秒钟
		<-time.Tick(5 * time.Second)
	}
}
