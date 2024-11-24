package main

import (
	"fmt"
	server "ratelimit/pkg"
	"ratelimit/pkg/config"
	"strconv"

	utils "ratelimit/utils"

	"github.com/caarlos0/env"
)

func main() {
	utils.InitLogger()
	cfg := config.InitConfig()
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)

	m, _ := strconv.Atoi(cfg.MaxRequest)
	b, _ := strconv.Atoi(cfg.BucketToken)
	s := server.New(m, b)
	s.InitRoutes()
	s.Run(":" + cfg.HttpPort)
}
