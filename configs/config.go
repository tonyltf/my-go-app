package config

import (
	"log"

	"github.com/spf13/viper"
)

type envConfigs struct {
	ApiInterval  int    `mapstructure:"API_INTERVAL"`
	ApiSource    string `mapstructure:"API_SOURCE"`
	DbConnection string `mapstructure:"DB_CONNECTION"`
	DbDriver     string `mapstructure:"DB_DRIVER"`
}

func InitConfig() (config *envConfigs) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("default")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return config
}
