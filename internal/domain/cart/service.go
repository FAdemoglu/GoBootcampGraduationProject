package cart

type CartService struct {
	r CartRepository
}

func NewCartService(r CartRepository) *CartService {
	return &CartService{
		r: r,
	}
}

func (r *CartService) GetAllCartProduct(pageIndex, pageSize int, customerUsername string) ([]Cart, int) {
	cartItems, count := r.r.GetAllCartProducts(pageIndex, pageSize, customerUsername)
	return cartItems, count
}

func (r *CartService) Create(c Cart) {
	r.r.AddToCart(c)
}
