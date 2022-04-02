package products

import "gorm.io/gorm"

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Migration() error {
	return r.db.AutoMigrate(&Product{})
}

func (r *ProductRepository) GetAllProducts(pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	var count int64
	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	return products, int(count)
}

func (r *ProductRepository) InserSampleData() {
	products := Product{
		Id:                1,
		ProductName:       "Lenovo IDEA PAD",
		ProductPrice:      300,
		ProductStockCount: 100,
		CategoryId:        1,
	}
	r.db.Where(Product{ProductName: products.ProductName}).Attrs(Product{Id: products.Id, ProductName: products.ProductName, ProductPrice: products.ProductPrice, ProductStockCount: products.ProductStockCount, CategoryId: products.CategoryId}).FirstOrCreate(&products)
}
