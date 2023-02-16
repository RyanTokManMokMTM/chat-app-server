package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire uint
	}

	MySQL struct {
		DataSource   string
		MaxIdleConns int
		MaxOpenConns int
	}
}
