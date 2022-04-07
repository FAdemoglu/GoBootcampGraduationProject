package order

import "errors"

var (
	ErrCouldNotFindOrderById = errors.New("Bu id ile bir sipariş bulunamadı")
	ErrLessThanZero          = errors.New("Girilen bazı değerler sıfırdan küçük")
	ErrCanNotCancelOrder     = errors.New("14 gün geçtiği için sipariş iptal edilemez")
)
