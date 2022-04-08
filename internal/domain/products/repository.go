package products

import (
	"errors"
	"gorm.io/gorm"
)

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

func (r *ProductRepository) DeleteProductById(id int) error {
	var exists bool
	result := r.db.Delete(&Product{}, id)

	if err := result.Scan(&exists); err != nil {
		return ErrCouldNotFindProductById
	} else if !exists {
		return ErrCouldNotFindProductById
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ProductRepository) CreateProduct(p Product) error {
	result := r.db.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ProductRepository) UpdateProduct(Id int, p Product) error {
	var product Product
	result := r.db.First(&product, Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ErrCouldNotFindProductById
	}
	if result.Error != nil {
		return result.Error
	}
	if p.ProductStockCount < 0 || p.CategoryId < 0 {
		return ErrLessThanZero
	}
	product.ProductStockCount = p.ProductStockCount
	product.ProductPrice = p.ProductPrice
	product.CategoryId = p.CategoryId
	p.ProductName = p.ProductName
	resultSave := r.db.Save(product)
	if resultSave.Error != nil {
		return resultSave.Error
	}
	return nil
}

func (r *ProductRepository) SearchProduct(pageIndex, pageSize int, searched string) ([]Product, int) {
	var products []Product
	var count int64
	r.db.Preload("Category").Joins("JOIN category on category.CategoryId=products.CategoryId").Where("ProductName LIKE ?", "%"+searched+"%").Or("category.CategoryName LIKE ?", "%"+searched+"%").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	return products, int(count)
}

func (r *ProductRepository) InserSampleData() {
	products := Product{
		Id:                1,
		ProductName:       "Lenovo IDEA PAD",
		ProductPrice:      300,
		ProductStockCount: 100,
		CategoryId:        1,
		ProductCode:       100001,
	}
	r.db.Where(Product{ProductName: products.ProductName}).Attrs(Product{Id: products.Id, ProductName: products.ProductName, ProductPrice: products.ProductPrice, ProductStockCount: products.ProductStockCount, CategoryId: products.CategoryId, ProductCode: products.ProductCode}).FirstOrCreate(&products)
}
