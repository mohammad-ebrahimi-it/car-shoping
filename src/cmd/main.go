package main

import (
	"github.com/mohammad-ebrahimi-it/car-shoping/api"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/cache"
)

func main() {
	cfg := config.GetConfig()
	api.InitServer(cfg)
	defer cache.CloseRedis()
	cache.InitRedis(cfg)
}
