package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func Logger(g *gin.Context) {
	file, err := os.OpenFile(time.Now().Format("2006-01-02"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		g.Next()
	}
	logger := log.New(file, time.Now().Format("2006-01-02"), os.O_CREATE|os.O_APPEND|os.O_WRONLY)
	logger.Printf("url is %s", g.Request.URL.Path)
	g.Next()
}
