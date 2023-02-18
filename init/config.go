package init

import (
	"log"

	"github.com/spf13/viper"
)

type envConfigs struct {
	ApiInterval int    `mapstructure:"API_INTERVAL"`
	ApiSource   string `mapstructure:"API_SOURCE"`
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
