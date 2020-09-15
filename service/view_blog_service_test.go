package service

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestAddViewTimes(t *testing.T) {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	viper.ReadInConfig()
	Init()
	fmt.Println(AddViewTimes("123456", "4566789"))
	fmt.Println(AddViewTimes("123456", "4566789"))
	fmt.Println(AddViewTimes("123456", "4566789"))
}
