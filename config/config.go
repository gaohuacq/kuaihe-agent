package config

type Config struct {
	Port                        string `yaml:"port"` // 端口
	ClientID                    string `yaml:"client_id"`
	ClientSecret                string `yaml:"client_secret"`
	EurekaAddress               string `yaml:"eureka_address"`                // 微服务列表获取地址
	AuthAddress                 string `yaml:"auth_address"`                  // 权限accesstoken获取地址
	ProcessAuthorizationSeconds int64  `yaml:"process_authorization_seconds"` // 提前处理accesstoken的秒数
	AuthorizationFresh          int64  `yaml:"authorization_fresh"`           // 刷新accesstoken的秒数

	Redis RedisConfig `yaml:"redis"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}
