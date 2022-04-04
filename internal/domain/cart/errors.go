package cart

import "errors"

var (
	ErrCouldNotFindCartById = errors.New("Bu id ile bir ürün bulunamadı")
	ErrLessThanZero         = errors.New("Girilen bazı değerler sıfırdan küçük")
)
