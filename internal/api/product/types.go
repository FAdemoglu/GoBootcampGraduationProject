package product

type ProductRequest struct {
	ProductName       string `json:"productname" binding:"required"`
	ProductPrice      int    `json:"productprice" binding:"required"`
	ProductStockCount int    `json:"stock" binding:"required"`
	CategoryId        int    `json:"categoryId" binding:"required"`
}
