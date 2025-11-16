package config

import (
	"github.com/yadav-shubh/base-backend/utils"
	"go.uber.org/zap"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	Server struct {
		Host string
	}
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		utils.Log.Error("Failed to load config", zap.Error(err))
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		utils.Log.Error("Failed to unmarshal config", zap.Error(err))
	}

	cfg = &c
}

func Get() *Config {
	once.Do(loadConfig)
	return cfg
}
