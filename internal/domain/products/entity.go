package products

import "github.com/FAdemoglu/graduationproject/internal/domain/category"

type Product struct {
	Id                int `gorm:"column:ProductID;autoIncrement;PRIMARY_KEY;not null"`
	ProductName       string
	ProductPrice      int
	ProductStockCount int
	CategoryId        int
	ProductCode       int
	Category          category.Category `gorm:"foreignKey:CategoryId;references:CategoryId" json:"Category,omitempty"`
}

func NewProduct(productName string, productPrice, productStockCount, CategoryId, ProductCode int) *Product {
	return &Product{
		ProductName:       productName,
		ProductPrice:      productPrice,
		ProductStockCount: productStockCount,
		CategoryId:        CategoryId,
		ProductCode:       ProductCode,
	}
}

func (Product) TableName() string {
	return "Products"
}
