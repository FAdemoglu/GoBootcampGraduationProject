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

// GetAllProducts GetCartAllProducts godoc
// @Summary Gets all cart products with paginated result
// @Tags Order
// @Accept  json
// @Produce  json
// @Param page query int false "Page Index"
// @Param pageSize query int false "Page Size"
// @Success 200 {object} pagination.Pages
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/list [get]
func (c *OrderController) GetAllProducts(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, "qa")
	items, count := c.orderService.GetAllOrderedProduct(pageIndex, pageSize, decodedClaims.Username)
	paginatedResult := pagination.NewFromGinRequest(g, count)
	paginatedResult.Items = items
	g.JSON(http.StatusOK, paginatedResult)
}

// CancelOrderById FromCart godoc
// @Summary Delete product from cart
// @Tags Order
// @Accept  json
// @Produce  json
// @Param OrderId query int false "OrderId"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /order/cancel [delete]
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

// CreateOrder godoc
// @Summary Add product to cart with token
// @Tags Order
// @Accept  json
// @Produce  json
// @Param CartId query int false "CartId"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /order/create [post]
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
