package util

import (
	"log"

	"github.com/spf13/viper"
)

type ProjectCongfig struct {
	DSN           string `mapstructure:"dsn"`
	ServerAddress string `mapstructure:"server_address"`
	EncKey        string `mapstructure:"encryption_key"`
}

var Config ProjectCongfig

func LoadConfig(path string) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatal(err)
	}
}
