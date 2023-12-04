package config

import (
	"errors"
	"fmt"
	"github.com/ArthurHlt/go-eureka-client/eureka"
)

var (
	// redis存储个服务 list存储 http://10.49.189.190:30600

	UserCenterAddress      = ""                       // 用户中心地址
	EvaluateCenterAddress  = ""                       // 评价中心
	ProductCenterAddress   = "PRODUCT-GATHER-GATEWAY" // 商品中心地址
	OrderCenterAddress     = ""                       // 订单中心
	CartCenterAddress      = ""                       // 购物车
	PromotionCenterAddress = ""                       // 促销中心
	DeliveryCenterAddress  = ""                       // 配送中心
	PayCenterAddress       = ""                       // 支付中心
	BasicServicesAddress   = ""                       // 基础服务地址
	PartnerCenterAddress   = ""                       // 门店中心
)

type EurekaEntity struct{}

func EurekaProviderServeAddress(address string) error {
	client := eureka.NewClient([]string{
		address, // 从Spring Boot基础的Eureka服务器获取地址
		// 添加其他Eureka服务器地址（如果有的话）
	})
	EurekaClient = client
	// 获取所有应用程序信息
	applications, err := client.GetApplications()
	if err != nil {
		return errors.New(fmt.Sprintf("无法获取应用程序信息：%v", err))
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
	return nil
}

func GetOneEurekaInstance(instanceId string) {

}
