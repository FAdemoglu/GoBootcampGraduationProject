package category

type Category struct {
	Id           int `gorm:"column:CategoryId;autoIncrement;PRIMARY_KEY;not null"`
	CategoryName string
}

func (Category) TableName() string {
	return "Category"
}
