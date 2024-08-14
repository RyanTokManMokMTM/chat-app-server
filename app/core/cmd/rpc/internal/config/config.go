package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	AssetsRPC zrpc.RpcClientConf
	MySQL     struct {
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
	RedisPubSub struct {
		Addr     string
		Password string
	}
}
