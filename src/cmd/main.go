package main

import (
	"github.com/mohammad-ebrahimi-it/car-shoping/api"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/cache"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/db"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/db/migrations"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	migrations.UP_1()

	defer db.CloseDb()

	api.InitServer(cfg)
}

//GOPROXY='https://goproxy.cn,https://proxy.golang.com.cn,direct'
