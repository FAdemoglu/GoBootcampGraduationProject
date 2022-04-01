package category

import "gorm.io/gorm"

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (c *CategoryRepository) Migration() error {
	return c.db.AutoMigrate(&Category{})
}

func (c *CategoryRepository) InsertSampleData() {
	categories := []*Category{
		{
			Id:           1,
			CategoryName: "Bilgisayar",
		},
		{
			Id:           2,
			CategoryName: "Tablet",
		},
	}

	for _, category := range categories {
		c.db.Where(Category{CategoryName: category.CategoryName}).Attrs(Category{Id: category.Id, CategoryName: category.CategoryName}).FirstOrCreate(&category)
	}
}
