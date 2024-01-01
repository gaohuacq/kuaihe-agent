package config

import (
	"fmt"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/go-redis/redis"
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

func InitRedis() error {
	server := fmt.Sprintf("%v:%v", GlobalConfig.Redis.Host, GlobalConfig.Redis.Port)
	c := redis.NewClient(&redis.Options{
		Addr:     server,
		Password: GlobalConfig.Redis.Password,
		DB:       0,
	})
	_, err := c.Ping().Result()
	if err != nil {
		return err
	}
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
