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

func (c *CategoryRepository) CreateCategoriesFromCsv(categories []Category) {
	for _, category := range categories {
		c.db.Where(Category{CategoryName: category.CategoryName}).Attrs(Category{Id: category.Id, CategoryName: category.CategoryName}).FirstOrCreate(&category)
	}
}

func (r *CategoryRepository) GetAllCategories(pageIndex, pageSize int) ([]Category, int) {
	var categories []Category
	var count int64

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count)
	return categories, int(count)
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
