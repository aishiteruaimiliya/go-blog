package main

import (
	"blog/config"
	"blog/model"
	"blog/routes"
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	service.Init()
	model.Init()
	fmt.Println(viper.GetString("server.addr"))
	g := gin.New()
	routes.InitRoutes(g)
	url:=fmt.Sprintf("%s:%d",viper.GetString("server.addr"),viper.GetInt("server.port"))
	err := g.Run(url)
	if err != nil {
		fmt.Println("start server err,err is ",err)
	}
}
