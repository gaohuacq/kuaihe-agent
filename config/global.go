package config

import (
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"runtime/debug"

	"github.com/coocood/freecache"
)

var (
	GlobalConfig *Config
	EurekaClient *eureka.Client
	FreeCache    *freecache.Cache
)

func InitFreeCache() error {
	cacheSize := 5 * 1024 * 1024 // 缓存5M 仅仅存储一个accesstoken
	cache := freecache.NewCache(cacheSize)
	debug.SetGCPercent(20)
	FreeCache = cache
	return nil
}

func InitConfig() error {
	yamlFiler, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	var configuration Config
	err = yaml.Unmarshal(yamlFiler, &configuration)
	if err != nil {
		return err
	}
	GlobalConfig = &configuration
	return nil
}
