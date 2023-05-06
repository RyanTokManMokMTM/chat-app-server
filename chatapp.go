package main

import (
	"flag"
	"fmt"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/config"
	_ "github.com/ryantokmanmokmtm/chat-app-server/internal/dao"
	routerConf "github.com/ryantokmanmokmtm/chat-app-server/router"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/chatapp.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	//fmt.Println(c.MaxFileSize)
	server := routerConf.ConfigRouter(c)

	fmt.Printf("Starting router at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
