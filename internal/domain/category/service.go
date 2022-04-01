package category

type CategoryService struct {
	r CategoryRepository
}

func NewCategoryService(c CategoryRepository) *CategoryService {
	return &CategoryService{
		r: c,
	}
}
