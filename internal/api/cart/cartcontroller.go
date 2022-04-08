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
func (c *CartController) GetAllProducts(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, os.Getenv("ENV"))
	items, count := c.cartService.GetAllCartProduct(pageIndex, pageSize, decodedClaims.Username)
	paginatedResult := pagination.NewFromGinRequest(g, count)
	paginatedResult.Items = items
	g.JSON(http.StatusOK, paginatedResult)
}

//Add to cart product with token
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
