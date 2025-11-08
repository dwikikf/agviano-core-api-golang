package config

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	AppName  string         `mapstructure:"APP_NAME"`
	AppEnv   string         `mapstructure:"APP_ENV"`
	AppPort  string         `mapstructure:"APP_PORT"`
	Database DatabaseConfig `mapstructure:",squash"`
	JWT      JWTConfig      `mapstructure:",squash"`
	LogLevel string         `mapstructure:"LOG_LEVEL"`
	Message  string         `mapstructure:"MESSAGE"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

type JWTConfig struct {
	SecretKey  string `mapstructure:"JWT_SECRET_KEY"`
	ExpireHour int    `mapstructure:"JWT_EXPIRE_HOUR"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetDefault("APP_NAME", "Agviano Core API")
	viper.SetDefault("APP_ENV", "Depelopment")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("MESSAGE", "Default")

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "")
	viper.SetDefault("DB_NAME", "agviano_core_db")

	viper.AddConfigPath(path)
	viper.SetConfigFile(filepath.Join(path, ".env"))

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Println("⚠️ .env file not found, relying on environment variables or defaults.")
		} else {
			return
		}
	} else {
		fmt.Println("✅ .env file loaded successfully.")
	}

	err = viper.Unmarshal(&config)
	return
}
