package order

import "time"

type Order struct {
	Id               int `gorm:"autoIncrement;primaryKey;not null"`
	OrderId          int `gorm:"column:OrderId;not null"`
	CustomerUsername string
	OrderItems       []OrderItem `gorm:"foreignKey:OrderId"`
	CreatedAt        time.Time   `json:"createdAt"`
}

type OrderItem struct {
	Id          int    `json:"id" gorm:"column:OrderItemId;autoIncrement;primaryKey;not null"`
	ProductName string `json:"productName"`
	UnitPrice   int    `json:"unitPrice"`
	Quantity    int    `json:"quantity"`
	ProductCode int    `json:"productCode"`
	ProductId   int    `json:"-"`
	OrderId     int    `json:"-"`
}

func (Order) TableName() string {
	return "Order"
}
