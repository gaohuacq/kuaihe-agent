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

	userCenterAddress      []string // 用户中心地址
	promotionCenterAddress []string // 促销中心
	partnerCenterAddress   []string // 门店中心
	productCenterAddress   []string // 商品中心

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
	addressList := strings.Split(address, ",")
	client := eureka.NewClient(addressList)
	EurekaClient = client
	// 获取所有应用程序信息
	applications, err := client.GetApplications()
	if err != nil {
		return errors.New(fmt.Sprintf("无法获取应用程序信息：%v", err))
	}

	ucenter := make([]string, 0)
	promotion := make([]string, 0)
	hzzOpenApi := make([]string, 0)

	// 遍历应用程序并打印它们的名称和实例数量
	for _, app := range applications.Applications {
		fmt.Printf("应用程序名称：%s\n", app.Name)
		fmt.Printf("实例数量：%d\n", len(app.Instances))
		fmt.Println("实例详细信息：")
		for _, instance := range app.Instances {
			// 券中心
			if strings.ToUpper(app.Name) == "PROMOTION-GATEWAY" {
				promotion = append(promotion, fmt.Sprintf("%v:%v", instance.HostName, instance.Port.Port))
			}

			if strings.ToUpper(app.Name) == "UCENTER-GATEWAY" {
				ucenter = append(ucenter, fmt.Sprintf("%v:%v", instance.HostName, instance.Port.Port))
			}

			if strings.ToUpper(app.Name) == "HZZ-OPENAPI-GATEWAY" {
				hzzOpenApi = append(hzzOpenApi, fmt.Sprintf("%v:%v", instance.HostName, instance.Port.Port))
			}

			if strings.ToUpper(app.Name) == "PRODUCT-GATHER-GATEWAY" {
				productCenterAddress = append(productCenterAddress, fmt.Sprintf("%v:%v", instance.HostName, instance.Port.Port))
			}
			fmt.Printf("  实例ID：%s\n", instance.InstanceID)
			fmt.Printf("  主机名：%s\n", instance.HostName)
			fmt.Printf("  地址：%s:%d\n", instance.HostName, instance.Port.Port)
			fmt.Printf("  健康状态：%s\n", instance.Status)
		}
	}

	if len(productCenterAddress) > 0 {
		setServiceAddressProductApi(productCenterAddress)
	} else {
		return errors.New("productcenter地址获取失败")
	}

	if len(ucenter) > 0 {
		setServiceAddressUCenter(ucenter)
	} else {
		return errors.New("ucenter地址获取失败")
	}
	if len(promotion) > 0 {
		setServiceAddressPromotion(promotion)
	} else {
		return errors.New("promotion服务地址获取失败")
	}
	if len(hzzOpenApi) > 0 {
		setServiceAddressHzzOpenApi(hzzOpenApi)
	} else {
		return errors.New("hzzOpenApi服务地址获取失败")
	}
	return nil
}

func setServiceAddressProductApi(productAddress []string) {
	mutex.Lock()
	defer mutex.Unlock()

	productCenterAddress = nil
	productCenterAddress = append(productCenterAddress, productAddress...)
}

func GetServiceAddressProductApi() []string {
	mutex.Lock()
	defer mutex.Unlock()

	return append([]string{}, productCenterAddress...)
}

func setServiceAddressHzzOpenApi(hzzOpenApi []string) {
	mutex.Lock()
	defer mutex.Unlock()

	partnerCenterAddress = nil
	partnerCenterAddress = append(partnerCenterAddress, hzzOpenApi...)
}

func GetServiceAddressHzzOpenApi() []string {
	mutex.Lock()
	defer mutex.Unlock()

	return append([]string{}, partnerCenterAddress...)
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
