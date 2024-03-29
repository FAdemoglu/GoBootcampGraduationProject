package order

type OrderService struct {
	r OrderRepository
}

func NewOrderService(r OrderRepository) *OrderService {
	return &OrderService{
		r: r,
	}
}

func (r *OrderService) GetAllOrderedProduct(pageIndex, pageSize int, customerUsername string) ([]Order, int) {
	orderItems, count := r.r.GetAllOrderedProducts(pageIndex, pageSize, customerUsername)
	return orderItems, count
}

func (r *OrderService) CancelOrder(orderId int, customerUsername string) error {
	error := r.r.CancelOrder(orderId, customerUsername)
	return error
}

func (r *OrderService) CreateOrder(order Order) error {
	err := r.r.CreateOrder(order)
	return err
}
