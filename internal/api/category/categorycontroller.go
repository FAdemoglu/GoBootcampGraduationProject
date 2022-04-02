package category

import (
	"fmt"
	"github.com/FAdemoglu/graduationproject/helper"
	"github.com/FAdemoglu/graduationproject/internal/domain/category"
	"github.com/FAdemoglu/graduationproject/pkg/pagination"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

type CategoryController struct {
	categoryService *category.CategoryService
}

func NewCategoryController(service *category.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: service,
	}
}

// GetAllCities godoc
// @Summary Gets all cities with paginated result
// @Tags City
// @Accept  json
// @Produce  json
// @Param page query int false "Page Index"
// @Param pageSize query int false "Page Size"
// @Success 200 {object} pagination.Pages
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /city [get]
func (c *CategoryController) GetAllCategories(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	items, count := c.categoryService.GetAllCategories(pageIndex, pageSize)
	paginatedResult := pagination.NewFromGinRequest(g, count)
	paginatedResult.Items = items

	g.JSON(http.StatusCreated, paginatedResult)
}

func (c *CategoryController) UploadCSVDatas(g *gin.Context) {
	file, header, err := g.Request.FormFile("categoriescsv")
	filename := header.Filename
	fmt.Println(filename)
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed to upload",
		})
	}

	// The file cannot be received.
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"content": file,
	})
	categories, _ := helper.ReadCsvToBookSlice(filename)
	c.categoryService.SaveCsvCategories(categories)
}
