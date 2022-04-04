package cart

import (
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) Migration() error {
	return r.db.AutoMigrate(&Cart{}, &Item{})
}

func (r *CartRepository) GetAllCartProducts(pageIndex, pageSize int, customerUsername string) ([]Cart, int) {
	var cart []Cart
	var count int64
	r.db.Preload("Items").Where("customerusername = ?", customerUsername).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&cart).Count(&count)
	return cart, int(count)
}

func (r *CartRepository) AddToCart(c Cart) error {
	result := r.db.Create(&c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CartRepository) InsertSampleData() {
	cart := Cart{
		CustomerUsername: "Furkan Ademoglu",
		Items: []Item{
			{
				ProductName: "Lenovo IDEA PAD",
				UnitPrice:   300,
				Quantity:    1,
			},
		},
	}

	r.db.Create(&cart)
}
