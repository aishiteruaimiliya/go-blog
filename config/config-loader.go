package config

import (
	"fmt"
	"github.com/spf13/viper"
)


func InitConfig(){
	fmt.Println("loading config..")
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("config load err.....",err)
	}
}
