package order

import (
	"fmt"
	"github.com/FAdemoglu/graduationproject/internal/config"
	"github.com/FAdemoglu/graduationproject/internal/domain/cart"
	"github.com/FAdemoglu/graduationproject/internal/domain/order"
	"github.com/FAdemoglu/graduationproject/internal/domain/products"
	jwtHelper "github.com/FAdemoglu/graduationproject/pkg/jwt"
	"github.com/FAdemoglu/graduationproject/pkg/pagination"
	"github.com/FAdemoglu/graduationproject/shared"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type OrderController struct {
	appConfig      *config.Configuration
	orderService   *order.OrderService
	productService *products.ProductService
	cartService    *cart.CartService
}

func NewOrderController(appConfig *config.Configuration, service *order.OrderService, productservice *products.ProductService, cartService *cart.CartService) *OrderController {
	return &OrderController{
		appConfig:      appConfig,
		orderService:   service,
		productService: productservice,
		cartService:    cartService,
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

func (c *OrderController) CancelOrderById(g *gin.Context) {
	IdForm := g.Query("OrderId")
	OrderId, _ := strconv.Atoi(IdForm)
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, "qa")
	error := c.orderService.CancelOrder(OrderId, decodedClaims.Username)
	if error == order.ErrCouldNotFindOrderById {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Bu id ile bir sipariş yoktur.",
		})
		return
	}
	if error == order.ErrCanNotCancelOrder {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "14 gün geçtiği için bu sipariş iptal edilemez",
		})
		return
	}
	data := shared.ApiResponse{IsSuccess: true, Message: "Başarılı"}
	g.JSON(http.StatusOK, data)

}

func (c *OrderController) CreateOrder(g *gin.Context) {
	IdForm := g.Query("CartId")
	CartId, _ := strconv.Atoi(IdForm)

	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, "qa")
	cart := c.cartService.GetByIdCart(decodedClaims.Username, CartId)
	Items := make([]order.OrderItem, 1)
	for i, c := range cart.Items {
		Items[i].OrderId = c.CartId
		Items[i].ProductName = c.ProductName
		Items[i].UnitPrice = c.UnitPrice
		Items[i].Quantity = c.Quantity
		Items[i].ProductCode = c.ProductCode
		Items[i].ProductId = c.ProductId
	}

	order := order.Order{
		CustomerUsername: cart.CustomerUsername,
		IsCancelled:      false,
		CreatedAt:        time.Now(),
		OrderId:          cart.CartId,
		OrderItems:       Items,
	}
	err := c.orderService.CreateOrder(order)
	fmt.Println(err)
}
