package products

type ProductService struct {
	r ProductRepository
}

func NewProductService(r ProductRepository) *ProductService {
	return &ProductService{
		r: r,
	}
}

func (r *ProductService) GetAllProducts(pageIndex, pageSize int) ([]Product, int) {
	products, count := r.r.GetAllProducts(pageIndex, pageSize)
	return products, count
}
