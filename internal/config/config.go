package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	MySQL struct {
		DataSource   string
		MaxIdleConns int
		MaxOpenConns int
	}

	Salt          string
	ResourcesPath string
	MaxFileSize   int64

	//RabbitMQ struct {
	//	DataSource string
	//}

	Redis struct {
		Addr     string
		Password string
	}
}
