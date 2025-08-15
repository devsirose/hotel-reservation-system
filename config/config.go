package config

import (
	"github.com/spf13/viper"
)

// Config store all configuration of the application
// The value is read by Viper from a config file or environment variables.
type Config struct {
	AppName        string `mapstructure:"APP_NAME"`
	DbDriver       string `mapstructure:"DB_DRIVER"`
	DbSource       string `mapstructure:"DB_SOURCE"`
	ServerHost     string `mapstructure:"SERVER_HOST"`
	HTTPServerPort string `mapstructure:"HTTP_SERVER_PORT"`
	GRPCServerPort string `mapstructure:"GRPC_SERVER_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
