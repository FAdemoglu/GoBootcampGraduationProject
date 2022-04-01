package category

type Category struct {
	Id           int `gorm:"column:CategoryId,primary_key;auto_increment;not_null"`
	CategoryName string
}

func (Category) TableName() string {
	return "Category"
}
