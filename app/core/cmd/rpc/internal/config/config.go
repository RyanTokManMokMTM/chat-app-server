package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MySQL struct {
		DataSource   string
		MaxIdleConns int
		MaxOpenConns int
	}
	Salt          string
	ResourcesPath string
	MaxFileSize   uint
	AuthConf      struct {
		AccessSecret string
		AccessExpire int64
	}
}
