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

func (r *ProductService) DeleteProductById(Id int) error {
	err := r.r.DeleteProductById(Id)
	return err
}

func (r *ProductService) CreateProduct(p Product) error {
	err := r.r.CreateProduct(p)
	return err
}

func (r *ProductService) UpdateProduct(Id int, p Product) error {
	err := r.r.UpdateProduct(Id, p)
	return err
}
