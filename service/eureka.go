package service

import (
	"fmt"
	"log"

	"github.com/ArthurHlt/go-eureka-client/eureka"
)

//http://10.49.189.190:30600/eureka/apps/PRODUCT-GATHER-GATEWAY
//Conf.ClientId = "kuaihe-mp-test"
//Conf.ClientSecret = "VMsULHIwjljb4x46FiTT6b75"
//Conf.GateApiAddress = "http://10.49.189.190:30600"

func EurekaAddress() {
	// 创建一个Eureka客户端
	client := eureka.NewClient([]string{
		"https://gateway-eureka-test.19k8s.cn/eureka", // 从Spring Boot基础的Eureka服务器获取地址
		// 添加其他Eureka服务器地址（如果有的话）
	})

	// 获取所有应用程序信息
	applications, err := client.GetApplications()
	if err != nil {
		log.Fatalf("无法获取应用程序信息：%v", err)
	}

	// 遍历应用程序并打印它们的名称和实例数量
	for _, app := range applications.Applications {
		fmt.Printf("应用程序名称：%s\n", app.Name)
		fmt.Printf("实例数量：%d\n", len(app.Instances))
		fmt.Println("实例详细信息：")
		for _, instance := range app.Instances {
			fmt.Printf("  实例ID：%s\n", instance.InstanceID)
			fmt.Printf("  主机名：%s\n", instance.HostName)
			fmt.Printf("  地址：%s:%d\n", instance.HostName, instance.Port.Port)
			fmt.Printf("  健康状态：%s\n", instance.Status)
		}
	}
}
