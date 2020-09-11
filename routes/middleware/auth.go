package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context){
	// todo 使用redis做验证
	//if !viper.GetBool("debug"){
	//	token:= c.Request.Header.Get("auth")
	//	if token != "iamtoken" {
	//		c.AbortWithStatusJSON(403, "未登录")
	//	}
	//}
	c.Next()

}
