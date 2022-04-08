package main

import (
	"fmt"
	"github.com/FAdemoglu/graduationproject/internal/api"
	"github.com/FAdemoglu/graduationproject/pkg/gracefulshut"
	"github.com/FAdemoglu/graduationproject/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"time"
)

// @title Gin City Service API
// @version 1.0
// @description City service api provides city informations.
// @termsOfService http://mywebsite.com/terms

// @contact.name API Support
// @contact.url http://mywebsite.com/support
// @contact.email support@mywebsite.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	os.Setenv("ENV", "qa")
	r := gin.Default()
	registerMiddlewares(r)
	api.RegisterHandlers(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("liten:%s\n", err)
		}
	}()

	gracefulshut.ShutdownGin(srv, time.Second*5)
}

func registerMiddlewares(r *gin.Engine) {
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)

	}))
	r.Use(gin.Recovery())
	r.Use(middleware.LatencyLoger())
}
