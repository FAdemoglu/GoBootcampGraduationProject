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

func (r *OrderRepository) CancelOrder(orderId int, customerUsername string) error {
	var order Order
	r.db.Where("OrderId = ? AND CustomerUsername=? AND IsCancelled=?", orderId, customerUsername, false).Find(&order)
	currentTime := time.Now()
	time := order.CreatedAt.AddDate(0, 0, 14)
	if order.CustomerUsername == "" {
		return ErrCouldNotFindOrderById
	}
	if currentTime.Before(time) {
		order.IsCancelled = true
		resultSave := r.db.Save(order)
		if resultSave.Error != nil {
			return resultSave.Error
		}
	} else {
		return ErrCanNotCancelOrder
	}
	return nil
}

func (r *OrderRepository) CreateOrder(o Order) error {
	result := r.db.Create(&o)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *OrderRepository) Migration() error {
	return r.db.AutoMigrate(&Order{}, &OrderItem{})
}

func (r *OrderRepository) InsertSampleData() {
	order := Order{
		CustomerUsername: "Furkan Ademoglu",
		OrderId:          1,
		CreatedAt:        time.Now(),
		IsCancelled:      false,
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
	orderTwo := Order{
		CustomerUsername: "Furkan Ademoglu",
		OrderId:          3,
		CreatedAt:        time.Now(),
		IsCancelled:      true,
		OrderItems: []OrderItem{
			{
				ProductName: "Lenovo IDEA PAD",
				UnitPrice:   300,
				Quantity:    1,
				ProductId:   5,
				ProductCode: 1010101,
				OrderId:     3,
			},
		},
	}

	r.db.Create(&order)
	r.db.Create(&orderTwo)
}
