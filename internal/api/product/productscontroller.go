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

//Get all product list with pagination
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

func (c *ProductController) CreateProduct(g *gin.Context) {
	var req ProductRequest

	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		return
	}
	product := products.Product{
		ProductName:       req.ProductName,
		ProductPrice:      req.ProductPrice,
		ProductStockCount: req.ProductStockCount,
		CategoryId:        req.CategoryId,
	}
	err := c.productService.CreateProduct(product)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"message": "Ürün başarılı bir şekilde eklendi",
	})
}

func (c *ProductController) UpdateProduct(g *gin.Context) {
	var req ProductRequest

	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		return
	}
	IdForm := g.Query("Id")
	if IdForm == "" {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		return
	}
	Id, _ := strconv.Atoi(IdForm)
	product := products.Product{
		ProductName:       req.ProductName,
		ProductPrice:      req.ProductPrice,
		ProductStockCount: req.ProductStockCount,
		CategoryId:        req.CategoryId,
	}
	err := c.productService.UpdateProduct(Id, product)
	if err == products.ErrCouldNotFindProductById {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Bu Id ile bir ürün bulunamadı",
		})
		return
	}
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Ürün başarılı bir şekilde güncellendi",
	})
}

func (c *ProductController) SearchProduct(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	searched := g.Query("searched")
	items, count := c.productService.SearchProduct(pageIndex, pageSize, searched)
	paginatedResult := pagination.NewFromGinRequest(g, count)
	paginatedResult.Items = items
	g.JSON(http.StatusOK, paginatedResult)
}
