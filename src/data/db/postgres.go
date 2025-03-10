package db

import (
	"fmt"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.SSlMode, cfg.Server.TimeZone)

	dbClient, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()

	err = sqlDb.Ping()

	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConnection)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConnections)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Println("Db connection established")

	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	conn, _ := dbClient.DB()
	conn.Close()
}
