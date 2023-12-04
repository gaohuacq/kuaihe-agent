package config

import (
	"fmt"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var (
	GlobalConfig *Config
	AccessToken  string
	TokenType    string
	RedisClient  *redis.Client
	EurekaClient *eureka.Client
	//GlobalRedis  *redis.c
)

func InitRedis() error {
	server := fmt.Sprintf("%v:%v", GlobalConfig.Redis.Host, GlobalConfig.Redis.Port)
	c := redis.NewClient(&redis.Options{
		Addr:     server,
		Password: GlobalConfig.Redis.Password,
	})
	_, err := c.Ping().Result()
	if err != nil {
		return err
	}
	RedisClient = c
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
