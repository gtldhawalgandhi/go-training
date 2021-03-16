package util

import (
	"github.com/spf13/viper"
)

var configV Config

// Config ...
type Config struct {
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	AppEnv        string `mapstructure:"APP_ENV"`
}

// LoadConfig ...
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("myapp")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err == nil {
		configV = config
	}
	return
}

func GetConfig() Config {
	return configV
}

func GetEnv() string {
	return configV.AppEnv
}
