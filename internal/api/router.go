package api

import (
	"github.com/FAdemoglu/graduationproject/internal/api/auth"
	cart2 "github.com/FAdemoglu/graduationproject/internal/api/cart"
	categoryController "github.com/FAdemoglu/graduationproject/internal/api/category"
	"github.com/FAdemoglu/graduationproject/internal/api/product"
	"github.com/FAdemoglu/graduationproject/internal/config"
	"github.com/FAdemoglu/graduationproject/internal/domain/cart"
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
	cartRepository := cart.NewCartRepository(db)

	//Migration is done
	categoryRepository.Migration()
	userRepository.Migration()
	productRepository.Migration()
	cartRepository.Migration()

	//Inserted sample datas to database
	userRepository.InsertSampleData()
	categoryRepository.InsertSampleData()
	productRepository.InserSampleData()
	//cartRepository.InsertSampleData()
	userService := users.NewUserService(*userRepository)
	categoryService := category.NewCategoryService(*categoryRepository)
	productService := products.NewProductService(*productRepository)
	cartService := cart.NewCartService(*cartRepository)
	categoryController := categoryController.NewCategoryController(categoryService)
	authController := auth.NewAuthController(AppConfig, *userService)
	productController := product.NewProductControler(productService)
	cartController := cart2.NewCartController(AppConfig, cartService, productService)

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
	productGroup.POST("/update", middleware.AuthMiddleware(AppConfig.JwtSettings.SecretKey), productController.UpdateProduct)
	productGroup.GET("/search", productController.SearchProduct)

	//Cart Group
	cartGroup := r.Group("/cart")
	cartGroup.GET("/list", middleware.AuthMiddlewareForCart(AppConfig.JwtSettings.SecretKey), cartController.GetAllProducts)
	cartGroup.POST("/add", middleware.AuthMiddlewareForCart(AppConfig.JwtSettings.SecretKey), cartController.AddToCart)
	cartGroup.POST("/remove", middleware.AuthMiddlewareForCart(AppConfig.JwtSettings.SecretKey), cartController.DeleteById)
}
