package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	EnvMode string `json:"env"`
	AppPort string `json:"port"`
	AppHost string `json:"host"`

	SecretKey string `json:"secret_key"`

	DB DBConnInfo `json:"db"`
}

type DBConnInfo struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	User        string `json:"user"`
	Pass        string `json:"password"`
	DBName      string `json:"name"`
	MaxOpenConn int    `json:"max_open_conn"`
	MaxIdleConn int    `json:"max_idle_conn"`
	MaxLifetime int    `json:"max_lifetime"`
}

var Config Configs
var EnvFile string

func LoadConfig() (conf Configs) {
	if EnvFile == "" {
		EnvFile = ".env"
	}

	if err := godotenv.Load(EnvFile); err != nil {
		Log.Info("No environtment %s file found, will take host variable instead!", EnvFile)
	}

	Config = Configs{
		EnvMode:   DefaultValueString("dev", os.Getenv("ENV_MODE")),
		AppPort:   DefaultValueString("8080", os.Getenv("APP_PORT")),
		AppHost:   DefaultValueString("localhost", os.Getenv("APP_HOST")),
		SecretKey: DefaultValueString("secret", os.Getenv("SECRET_KEY")),
		DB: DBConnInfo{
			Host:        DefaultValueString("localhost", os.Getenv("DB_HOST")),
			Port:        DefaultValueString("5432", os.Getenv("DB_PORT")),
			User:        DefaultValueString("postgres", os.Getenv("DB_USER")),
			Pass:        DefaultValueString("postgres", os.Getenv("DB_PASSWORD")),
			DBName:      DefaultValueString("postgres", os.Getenv("DB_NAME")),
			MaxOpenConn: DefaultValueIntFromString(10, os.Getenv("DB_MAX_OPEN_CONN")),
			MaxIdleConn: DefaultValueIntFromString(5, os.Getenv("DB_MAX_IDLE_CONN")),
			MaxLifetime: DefaultValueIntFromString(5, os.Getenv("DB_MAX_LIFETIME")),
		},
	}

	return
}
