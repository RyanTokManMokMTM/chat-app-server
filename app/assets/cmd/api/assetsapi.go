package main

import (
	"api/app/assets/cmd/api/internal/router"
	"flag"
	"fmt"

	"api/app/assets/cmd/api/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/assetsapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := router.ConfigRouter(c)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
