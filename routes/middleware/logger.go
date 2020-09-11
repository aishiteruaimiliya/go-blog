package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger(g *gin.Context){
	fmt.Printf("receive req at %s,url is %s\n",time.Now(),g.Request.URL)
	g.Next()
}
