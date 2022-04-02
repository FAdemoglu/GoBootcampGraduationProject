package api

import (
	"github.com/FAdemoglu/graduationproject/internal/api/auth"
	categoryController "github.com/FAdemoglu/graduationproject/internal/api/category"
	"github.com/FAdemoglu/graduationproject/internal/api/product"
	"github.com/FAdemoglu/graduationproject/internal/config"
	"github.com/FAdemoglu/graduationproject/internal/domain/category"
	"github.com/FAdemoglu/graduationproject/internal/domain/products"
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

	//New Repositories created
	userRepository := users.NewUserRepository(db)
	categoryRepository := category.NewCategoryRepository(db)
	productRepository := products.NewProductRepository(db)

	//Migration is done
	categoryRepository.Migration()
	userRepository.Migration()
	productRepository.Migration()

	//Inserted sample datas to database
	userRepository.InsertSampleData()
	categoryRepository.InsertSampleData()
	productRepository.InserSampleData()
	userService := users.NewUserService(*userRepository)
	categoryService := category.NewCategoryService(*categoryRepository)
	productService := products.NewProductService(*productRepository)
	categoryController := categoryController.NewCategoryController(categoryService)
	authController := auth.NewAuthController(AppConfig, *userService)
	productController := product.NewProductControler(productService)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//User Group
	authGroup := r.Group("/user")
	authGroup.POST("/login", authController.Login)
	authGroup.POST("/register", authController.Register)
	authGroup.GET("/decode", middleware.AuthMiddleware(AppConfig.JwtSettings.SecretKey), middleware.AuthMiddleware(AppConfig.JwtSettings.SecretKey), authController.VerifyToken)

	//Categories Group
	categoryGroup := r.Group("/category")
	categoryGroup.GET("/list", categoryController.GetAllCategories)
	categoryGroup.POST("/uploadcsv", middleware.AuthMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.UploadCSVDatas)

	//Product Group
	productGroup := r.Group("/product")
	productGroup.GET("/list", productController.GetAllProducts)
	productGroup.DELETE("/remove", middleware.AuthMiddleware(AppConfig.JwtSettings.SecretKey), productController.DeleteProductById)
	productGroup.POST("/create", middleware.AuthMiddleware(AppConfig.JwtSettings.SecretKey), productController.CreateProduct)

}
