package order

import (
	"github.com/FAdemoglu/graduationproject/internal/config"
	"github.com/FAdemoglu/graduationproject/internal/domain/order"
	"github.com/FAdemoglu/graduationproject/internal/domain/products"
	jwtHelper "github.com/FAdemoglu/graduationproject/pkg/jwt"
	"github.com/FAdemoglu/graduationproject/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
	appConfig      *config.Configuration
	orderService   *order.OrderService
	productService *products.ProductService
}

func NewOrderController(appConfig *config.Configuration, service *order.OrderService, productservice *products.ProductService) *OrderController {
	return &OrderController{
		appConfig:      appConfig,
		orderService:   service,
		productService: productservice,
	}
}

func (c *OrderController) GetAllProducts(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, "qa")
	items, count := c.orderService.GetAllOrderedProduct(pageIndex, pageSize, decodedClaims.Username)
	paginatedResult := pagination.NewFromGinRequest(g, count)
	paginatedResult.Items = items
	g.JSON(http.StatusOK, paginatedResult)
}
