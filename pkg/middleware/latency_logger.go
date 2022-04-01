package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LatencyLoger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()

		context.Set("token", context.GetHeader("Authorization"))

		log.Println("This is before request logging")

		latency := time.Since(t)
		log.Println(latency)

		status := context.Writer.Status()
		log.Println(status)
	}
}
