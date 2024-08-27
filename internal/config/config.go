package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	Server      ServerConfig
	Environment string
	Google      GoogleConfig
	SerpAPI     API
	Secrets     SecretConfigs
}

type SecretConfigs struct {
	SerpAPIKey string `json:"serpAPIKey"`
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

func (c *Config) IsProductionEnv() bool {
	return c.Environment == "production"
}

func InitConfig(configname string) *Config {
	viper.SetConfigName(configname)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
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

	err := replaceEnvVariablesInViper()
	if err != nil {
		fmt.Println("Error processing the YAML file:", err)
		panic(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("fatal error unable to unmarshal config file: %s", err))
	}

	return &AppConfig
}

func replaceEnvVariablesInViper() error {
	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)

		// Check if value contains a placeholder for an environment variable
		if strings.Contains(value, "${") {
			// Extract the environment variable from the placeholder
			envVar := extractEnvVar(value)
			if envVar != "" {
				envValue := os.Getenv(envVar)
				if envValue != "" {
					// Replace the placeholder with the actual environment variable value
					viper.Set(key, envValue)
				} else {
					fmt.Printf("Environment variable %s is not set\n", envVar)
				}
			}
		}
	}
	return nil
}

// extractEnvVar extracts the environment variable name from the placeholder
func extractEnvVar(value string) string {
	start := strings.Index(value, "${")
	end := strings.Index(value, "}")
	if start != -1 && end != -1 && end > start {
		return value[start+2 : end]
	}
	return ""
}

func GetConfig() *Config {
	return &AppConfig
}

func (c *Config) ListenAddress() string {
	return ":" + strconv.Itoa(c.Server.Port)
}

func (c *Config) GetSerpAPI() string {
	return c.SerpAPI.Path
}

func (c *Config) GetSerpAPITimeOutInMs() time.Duration {
	duration, err := time.ParseDuration(strconv.Itoa(c.SerpAPI.TimeoutInMs) + "ms")
	if err != nil {
		panic(err)
	}
	return duration
}

func (c *Config) GetSecrets() map[string]string {
	mp := make(map[string]string)
	mp["serpAPIKey"] = c.Secrets.SerpAPIKey
	return mp
}
