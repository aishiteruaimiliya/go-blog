package middleware

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	// todo 使用redis做验证
	token := c.Request.Header.Get("token")
	AccountID, _ := service.RedisToToken(token)
	if token == "" {
		if c.Request.URL.Path == "/account/login" || c.Request.URL.Path == "/account/register" {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "未登录，请登录")
		}

	} else {
		c.Set("aid", AccountID)
		c.Next()
	}

}
