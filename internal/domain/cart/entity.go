package cart

type Cart struct {
	CartId           int `gorm:"column:CartId;autoIncrement;primaryKey;not null"`
	CustomerUsername string
	Items            []Item `gorm:"foreignKey:CartId"`
}

type Item struct {
	Id          int    `json:"id" gorm:"column:ItemId;autoIncrement;primaryKey;not null"`
	ProductName string `json:"productName"`
	UnitPrice   int    `json:"unitPrice"`
	Quantity    int    `json:"quantity"`
	ProductId   int    `json:"-"`
	CartId      int    `json:"-"`
}

func (Cart) TableName() string {
	return "Cart"
}
