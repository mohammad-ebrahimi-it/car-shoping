package migrations

import (
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/db"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/models"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func UP_1() {
	database := db.GetDb()

	var tables []interface{}

	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	err := database.Migrator().CreateTable(tables...)

	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}

	logger.Info(logging.Postgres, logging.Migration, "migration succeeded", nil)
}
