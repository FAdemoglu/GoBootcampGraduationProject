package cart

import (
	"github.com/FAdemoglu/graduationproject/internal/config"
	"github.com/FAdemoglu/graduationproject/internal/domain/cart"
	"github.com/FAdemoglu/graduationproject/internal/domain/products"
	jwtHelper "github.com/FAdemoglu/graduationproject/pkg/jwt"
	"github.com/FAdemoglu/graduationproject/pkg/pagination"
	"github.com/FAdemoglu/graduationproject/shared"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type CartController struct {
	appConfig      *config.Configuration
	cartService    *cart.CartService
	productService *products.ProductService
}

func NewCartController(appConfig *config.Configuration, service *cart.CartService, productservice *products.ProductService) *CartController {
	return &CartController{
		appConfig:      appConfig,
		cartService:    service,
		productService: productservice,
	}
}

//Get all products with database

// GetAllProducts GetCartAllProducts godoc
// @Summary Gets all cart products with paginated result
// @Tags Cart
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
func (c *CartController) GetAllProducts(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, os.Getenv("ENV"))
	items, count := c.cartService.GetAllCartProduct(pageIndex, pageSize, decodedClaims.Username)
	paginatedResult := pagination.NewFromGinRequest(g, count)
	paginatedResult.Items = items
	g.JSON(http.StatusOK, paginatedResult)
}

//Add to cart product with token

// AddToCart godoc
// @Summary Add product to cart with token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param cartAddRequest body CartAddRequest true "cart informations"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/add [post]
func (c *CartController) AddToCart(g *gin.Context) {
	var req CartAddRequest
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		return
	}
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, "qa")

	cart := cart.Cart{
		CustomerUsername: decodedClaims.Username,
		Items: []cart.Item{
			{
				ProductName: req.ProductName,
				UnitPrice:   req.UnitPrice,
				Quantity:    req.Quantity,
				ProductId:   req.ProductId,
			},
		},
	}
	c.cartService.Create(cart)
	data := shared.ApiResponse{IsSuccess: true, Message: "Success"}
	g.JSON(http.StatusOK, data)
}

//Delete product by Id

// DeleteById FromCart godoc
// @Summary Delete product from cart
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param Id query int false "Id"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/remove [delete]
func (c *CartController) DeleteById(g *gin.Context) {
	IdForm := g.Query("Id")
	Id, _ := strconv.Atoi(IdForm)
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, "qa")
	err := c.cartService.DeleteById(decodedClaims.Username, Id)

	if err == cart.ErrCouldNotFindCartById {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "There is no product with this Id.",
		})
		return
	}

	data := shared.ApiResponse{IsSuccess: true, Message: "Success"}
	g.JSON(http.StatusOK, data)
}

//Update cart with given parameter

// UpdateCart  godoc
// @Summary Add product to cart with token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param Id query int false "Id"
// @Param ItemId query int false "ItemId"
// @Param Count query int false "Count"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/update [put]
func (c *CartController) UpdateCart(g *gin.Context) {
	IdForm := g.Query("Id")
	ItemIdForm := g.Query("ItemId")
	CountForm := g.Query("Count")
	Id, _ := strconv.Atoi(IdForm)
	ItemId, _ := strconv.Atoi(ItemIdForm)
	Count, _ := strconv.Atoi(CountForm)
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, os.Getenv("ENV"))

	//Update cart if err exist return the error to the client
	err := c.cartService.UpdateTheCart(decodedClaims.Username, Id, ItemId, Count)
	if err == cart.ErrCouldNotFindCartById {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "There is no product with this id.",
		})
		return
	}

	data := shared.ApiResponse{IsSuccess: true, Message: "Success"}
	g.JSON(http.StatusOK, data)

}
