package category

type CategoryService struct {
	r CategoryRepository
}

func NewCategoryService(c CategoryRepository) *CategoryService {
	return &CategoryService{
		r: c,
	}
}

func (c *CategoryService) GetAllCategories(pageIndex, pageSize int) ([]Category, int) {
	categories, count := c.r.GetAllCategories(pageIndex, pageSize)
	return categories, count
}

func (c *CategoryService) SaveCsvCategories(category []Category) {
	c.r.CreateCategoriesFromCsv(category)
}
