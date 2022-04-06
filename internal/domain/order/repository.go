package order

import (
	"gorm.io/gorm"
	"time"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetAllOrderedProducts(pageIndex, pageSize int, customerUsername string) ([]Order, int) {
	var cart []Order
	var count int64
	r.db.Preload("OrderItems").Where("CustomerUsername = ?", customerUsername).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&cart).Count(&count)
	return cart, int(count)
}

func (r *OrderRepository) Migration() error {
	return r.db.AutoMigrate(&Order{}, &OrderItem{})
}

func (r *OrderRepository) InsertSampleData() {
	order := Order{
		CustomerUsername: "Furkan Ademoglu",
		OrderId:          1,
		CreatedAt:        time.Now(),
		OrderItems: []OrderItem{
			{
				ProductName: "Lenovo IDEA PAD",
				UnitPrice:   300,
				Quantity:    1,
				ProductId:   5,
				ProductCode: 1010101,
				OrderId:     1,
			},
		},
	}

	r.db.Create(&order)
}
