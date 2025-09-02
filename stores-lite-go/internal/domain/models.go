package domain

type Product struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	PriceCents int64  `json:"priceCents"`
	Stock      int64  `json:"stock"`
}

type Order struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"productId"`
	Qty       int64 `json:"qty"`
	Amount    int64 `json:"amount"`
}
