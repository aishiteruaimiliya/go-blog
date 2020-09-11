package middleware

import (

	"github.com/gin-gonic/gin"
	"strings"
)

func RoleControl(c *gin.Context){
	url:=c.Request.URL
	if strings.Index(url.String(),"/admin")>=0{
	}
}
