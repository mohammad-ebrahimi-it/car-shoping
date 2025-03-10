package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
}
type ServerConfig struct {
	Port     string
	RunMode  string
	TimeZone string
}

type PostgresConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	DbName             string
	SSlMode            string
	MaxIdleConnection  int
	MaxOpenConnections int
	ConnMaxLifetime    time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	Db                 string
	Password           string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
	IdleCheckFrequency time.Duration
}

type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

func GetConfig() *Config {
	cfgConfig := getConfigPath("APP_ENV")
	v, err := LoadConfig(cfgConfig, "yml")

	if err != nil {
		log.Fatalf("LoadConfig failed: %v", err)
	}

	cfg, err := ParsConfig(v)

	if err != nil {
		log.Fatalf("ParsConfig failed: %v", err)
	}

	return cfg

}

func ParsConfig(v *viper.Viper) (*Config, error) {
	var cfg Config

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)

		return nil, err
	}

	return &cfg, nil
}

func LoadConfig(fileName, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	err := v.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "config/config-docker"
	} else if env == "development" {
		return "config/config-production"
	} else {
		return "../config/config-development"
	}
}
