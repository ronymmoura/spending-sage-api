package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment      string `mapstructure:"ENVIRONMENT"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabasePort     int    `mapstructure:"DATABASE_PORT"`

	ClerkKey string `mapstructure:"CLERK_KEY"`

	CacheUrl      string `mapstructure:"CACHE_URL"`
	CachePassword string `mapstructure:"CACHE_PASSWORD"`
	CacheDatabase int    `mapstructure:"CACHE_DATABASE"`

	DatabaseUrl string
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	config.DatabaseUrl = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.DatabaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)

	return
}
