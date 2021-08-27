package dto

type Order struct {
	OrderId   string
	UserId    string
	ProductId string
	Qty       int
	Amount    float64
	Status    string
}
