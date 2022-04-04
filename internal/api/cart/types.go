package cart

type CartAddRequest struct {
	ProductName string `json:"productname" binding:"required" `
	UnitPrice   int    `json:"unitprice" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required" `
}
