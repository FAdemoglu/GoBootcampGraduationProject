package products

import "errors"

var (
	ErrCouldNotFindProductById = errors.New("Bu id ile bir ürün bulunamadı")
	ErrLessThanZero            = errors.New("Girilen bazı değerler sıfırdan küçük")
)
