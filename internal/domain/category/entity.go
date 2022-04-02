package category

type Category struct {
	Id           int `gorm:"column:CategoryId;autoIncrement;primaryKey;not null"`
	CategoryName string
}

func (Category) TableName() string {
	return "Category"
}
