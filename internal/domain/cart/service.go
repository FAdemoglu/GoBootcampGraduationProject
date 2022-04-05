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

func (r *CartService) DeleteById(username string, id int) error {
	err := r.r.DeleteById(username, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *CartService) UpdateTheCart(username string, id, itemId, count int) error {
	err := r.r.UpdateTheCart(username, id, itemId, count)
	if err != nil {
		return err
	}
	return nil
}
