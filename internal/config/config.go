package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	Server      ServerConfig
	Environment string
	Google      GoogleConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type GoogleConfig struct {
	GoogleLensAPI API
}

type API struct {
	Path        string
	TimeoutInMs int `mapstructure:"timeoutInMs"`
}

func InitDefaultConfig() *Config {
	return InitConfig("application")
}

func InitConfig(configname string) *Config {
	viper.AutomaticEnv()
	viper.SetConfigName(configname)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("../../config/")
	viper.AddConfigPath("../../../config/")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Cannot read config file config/application.yaml")
		panic(err)
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("fatal error unable to unmarshal config file: %s", err))
	}

	return &AppConfig
}

func GetConfig() *Config {
	return &AppConfig
}

func (c *Config) ListenAddress() string {
	return ":" + strconv.Itoa(c.Server.Port)
}
