package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"product_kuaihe/config"
	"product_kuaihe/model"
	modelPro "product_kuaihe/model/product_center"
	"product_kuaihe/service/product_center"
	"strings"
)

func main() {
	// 配置初始化
	if err := config.InitConfig(); err != nil {
		log.Fatal("yaml配置文件读取失败")
		return
	}

	// redis初始化
	if err := config.InitRedis(); err != nil {
		log.Fatal("redis初始化失败")
		return
	}

	// eureka服务发现初始化到redis
	if err := config.EurekaProviderServeAddress(config.GlobalConfig.EurekaAddress); err != nil {
		log.Fatal(err)
		return
	}

	// 创建路由
	router := fasthttprouter.New()
	router.POST("/product/search", ProductSearch)
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%v", config.GlobalConfig.Port), router.Handler))
}

func ProductSearch(ctx *fasthttp.RequestCtx) {
	resp, err := product_center.ProductSearch(modelPro.ProductSearchReq{
		Channel:      "chihe",
		Size:         5,
		ProvinceCode: "510000",
		Keyword:      "五粮液",
		SortType:     "PRICE_ASC",
		Accross:      true,
		LocateLat:    "30.567067",
		LocateLon:    "104.064753",
		Page:         1,
	})
	if err != nil {
		fmt.Println(err)
	}
	bJson, _ := json.Marshal(resp)
	ctx.Response.Write(bufio.NewWriter(bytes.NewBuffer(bJson)))
}

// Auth 测试方便用 获取clientid的 accesstoken
func Auth() {
	postData := url.Values{}
	postData.Add("grant_type", "client_credentials")
	postData.Add("client_id", config.GlobalConfig.ClientID)
	postData.Add("client_secret", config.GlobalConfig.ClientSecret)
	response, err := http.Post(config.GlobalConfig.AuthAddress+"/oauth/token", "application/x-www-form-urlencoded", strings.NewReader(postData.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	bJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var authResponse model.AuthorizationResp
	err = json.Unmarshal(bJson, &authResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("authinfo: %+v\n", authResponse)
}
