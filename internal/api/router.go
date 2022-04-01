package api

import (
	"github.com/FAdemoglu/graduationproject/internal/api/auth"
	"github.com/FAdemoglu/graduationproject/internal/config"
	"github.com/FAdemoglu/graduationproject/internal/domain/category"
	"github.com/FAdemoglu/graduationproject/internal/domain/users"
	"github.com/FAdemoglu/graduationproject/pkg/database_handler"
	"github.com/FAdemoglu/graduationproject/pkg/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

var AppConfig = &config.Configuration{}

func RegisterHandlers(r *gin.Engine) {
	cfgFile := "./config/location.qa" + ".yaml"
	AppConfig, err := config.GetAllConfigValues(cfgFile)
	if err != nil {
		log.Fatalf("Failed to read config file. %v", err.Error())
	}

	db := database_handler.NewMySQLDB(AppConfig.DatabaseSettings.DatabaseURI)
	userRepository := users.NewUserRepository(db)
	categoryRepository := category.NewCategoryRepository(db)
	categoryRepository.Migration()
	userRepository.Migration()
	userRepository.InsertSampleData()
	categoryRepository.InsertSampleData()
	//userService := users.NewUserService(*userRepository)
	authController := auth.NewAuthController(AppConfig, *userRepository)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authGroup := r.Group("/user")
	authGroup.POST("/login", authController.Login)
	authGroup.GET("/decode", middleware.AuthMiddleware(AppConfig.JwtSettings.SecretKey), middleware.AuthMiddleware(AppConfig.JwtSettings.SecretKey), authController.VerifyToken)
}
