package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	HTTPServerAddress          string `mapstructure:"HTTP_SERVER_ADDRESS"`
	ElasticsearchServerAddress string `mapstructure:"ELASTICSEARCH_SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
