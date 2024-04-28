package util

import (
	"log"

	"github.com/spf13/viper"
)

type ProjectCongfig struct {
	DSN           string `mapstructure:"DSN"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	EncKey        string `mapstructure:"ENCRYPTION_KEY"`
}

var Config ProjectCongfig

func LoadConfig(path string) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatal(err)
	}
}
