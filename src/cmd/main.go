package main

import (
	"github.com/mohammad-ebrahimi-it/car-shoping/api"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/cache"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/db"
)

func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	if err != nil {
		panic(err)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil {
		panic(err)
	}
	db.CloseDb()

	api.InitServer(cfg)
}

//GOPROXY='https://goproxy.cn,https://proxy.golang.com.cn,direct'
