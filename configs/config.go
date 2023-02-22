package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type envConfigs struct {
	ApiInterval  int    `mapstructure:"API_INTERVAL"`
	ApiSource    string `mapstructure:"API_SOURCE"`
	DbConnection string `mapstructure:"DB_CONNECTION"`
	DbDriver     string `mapstructure:"DB_DRIVER"`
}

func InitConfig() (config *envConfigs) {

	_, b, _, _ := runtime.Caller(0)
	configpath := filepath.Dir(b)
	viper.AddConfigPath(configpath)
	viper.AddConfigPath("./")
	viper.AddConfigPath("/app")
	viper.AddConfigPath("/app/configs")
	viper.AddConfigPath("/my-go-app/configs")
	viper.SetConfigName("default.json")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file: ", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return config
}
