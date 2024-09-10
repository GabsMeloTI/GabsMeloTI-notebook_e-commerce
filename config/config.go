package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
	SSLMode  string
}

func init() {
	viper.SetDefault("api.port", "8000")
	viper.SetDefault("database.sslmode", "prefer")

	viper.AutomaticEnv()
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType(".env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
		log.Println("Configuration file not found, using environment variables and default values.")
	}

	cfg = &Config{
		API: APIConfig{
			Port: viper.GetString("api.port"),
		},
		DB: DBConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("database.user"),
			Pass:     viper.GetString("database.pass"),
			Database: viper.GetString("database.database"),
			SSLMode:  viper.GetString("database.sslmode"),
		},
	}

	if err := validateConfig(cfg); err != nil {
		return err
	}

	return nil
}

func validateConfig(cfg *Config) error {
	if cfg.DB.Host == "" || cfg.DB.Port == "" || cfg.DB.User == "" || cfg.DB.Pass == "" || cfg.DB.Database == "" {
		return logError("Database configuration is incomplete")
	}

	if cfg.API.Port == "" {
		return logError("API port configuration is missing")
	}

	return nil
}

func logError(msg string) error {
	err := log.New(os.Stderr, "CONFIG ERROR: ", log.LstdFlags)
	err.Println(msg)
	return msg
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
