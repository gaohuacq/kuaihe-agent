package config

import (
	"errors"
	"fmt"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"strings"
	"sync"
)

var (
	mutex sync.Mutex

	userCenterAddress      []string // "http://10.49.189.147:80"  // 用户中心地址
	promotionCenterAddress []string // "http://10.49.170.2:31559" // 促销中心

	EvaluateCenterAddress = "" // 评价中心
	ProductCenterAddress  = "" // 商品中心地址
	OrderCenterAddress    = "" // 订单中心
	CartCenterAddress     = "" // 购物车
	DeliveryCenterAddress = "" // 配送中心
	PayCenterAddress      = "" // 支付中心
	BasicServicesAddress  = "" // 基础服务地址
	PartnerCenterAddress  = "" // 门店中心

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

	ucenter := make([]string, 0)
	promotion := make([]string, 0)

	// 遍历应用程序并打印它们的名称和实例数量
	for _, app := range applications.Applications {
		//fmt.Printf("应用程序名称：%s\n", app.Name)
		//fmt.Printf("实例数量：%d\n", len(app.Instances))
		//fmt.Println("实例详细信息：")
		for _, instance := range app.Instances {
			// 券中心
			if strings.ToUpper(app.Name) == "PROMOTION-GATEWAY" {
				promotion = append(promotion, fmt.Sprintf("%v:%v", instance.HostName, instance.Port.Port))
			}

			if strings.ToUpper(app.Name) == "UCENTER-GATEWAY" {
				ucenter = append(ucenter, fmt.Sprintf("%v:%v", instance.HostName, instance.Port.Port))
			}
			//fmt.Printf("  实例ID：%s\n", instance.InstanceID)
			//fmt.Printf("  主机名：%s\n", instance.HostName)
			//fmt.Printf("  地址：%s:%d\n", instance.HostName, instance.Port.Port)
			//fmt.Printf("  健康状态：%s\n", instance.Status)
		}
	}
	setServiceAddressUCenter(ucenter)
	setServiceAddressPromotion(promotion)
	return nil
}

func setServiceAddressUCenter(address []string) {
	mutex.Lock()
	defer mutex.Unlock()

	userCenterAddress = nil
	userCenterAddress = append(userCenterAddress, address...)
}

func GetServiceAddressUCenter() []string {
	mutex.Lock()
	defer mutex.Unlock()

	return append([]string{}, userCenterAddress...)
}

func setServiceAddressPromotion(address []string) {
	mutex.Lock()
	defer mutex.Unlock()

	promotionCenterAddress = nil
	promotionCenterAddress = append(promotionCenterAddress, address...)
}

func GetServiceAddressPromotion() []string {
	mutex.Lock()
	defer mutex.Unlock()

	return append([]string{}, promotionCenterAddress...)
}
