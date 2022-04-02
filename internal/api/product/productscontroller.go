package product

import (
	"github.com/FAdemoglu/graduationproject/internal/domain/products"
	"github.com/FAdemoglu/graduationproject/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService *products.ProductService
}

func NewProductControler(service *products.ProductService) *ProductController {
	return &ProductController{
		productService: service,
	}
}

func (c *ProductController) GetAllProducts(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	items, count := c.productService.GetAllProducts(pageIndex, pageSize)
	paginatedResult := pagination.NewFromGinRequest(g, count)
	paginatedResult.Items = items
	g.JSON(http.StatusOK, paginatedResult)
}

func (c *ProductController) DeleteProductById(g *gin.Context) {
	IdForm := g.Query("Id")
	Id, _ := strconv.Atoi(IdForm)
	err := c.productService.DeleteProductById(Id)
	if err == products.ErrCouldNotFindProductById {
		g.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Bu id ile bir ürün bulunamadı",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"message": "Ürün başarılı bir şekilde silindi",
	})
}
