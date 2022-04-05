package cart

import (
	"errors"
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

func (r *CartRepository) UpdateTheCart(username string, id, itemId, count int) error {
	var cart Cart
	result := r.db.Preload("Items").Joins("JOIN item on item.CartId=cart.CartId").Where("CustomerUsername = ?", username).First(&cart, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ErrCouldNotFindCartById
	}
	if result.Error != nil {
		return result.Error
	}
	for i := range cart.Items {
		if cart.Items[i] == cart.Items[itemId-1] {
			cart.Items[i].Quantity += count
		}
	}
	resultSave := r.db.Save(&cart)
	if resultSave.Error != nil {
		return resultSave.Error
	}
	return nil
}

func (r *CartRepository) DeleteById(username string, id int) error {
	var exists bool
	result := r.db.Where("CustomerUsername = ?", username).Delete(&Cart{}, id)
	if err := result.Scan(&exists); err != nil {
		return result.Error
	} else if !exists {
		return ErrCouldNotFindCartById
	}
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
				ProductId:   5,
			},
		},
	}

	r.db.Create(&cart)
}
