package types

type OrderItem struct {
	ProductName string  `json:"product_name"`
	Picture     string  `json:"picture"`
	Qty         uint32  `json:"qty"`
	Cost        float32 `json:"cost"`
}

type Order struct {
	OrderId     string      `json:"order_id"`
	CreatedDate string      `json:"create_date"`
	Cost        float32     `json:"cost"`
	Items       []OrderItem `json:"items"`
}
