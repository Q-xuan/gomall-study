package types

type OrderItem struct {
	ProductID string  `json:"product_id"`
	Picture   string  `json:"picture"`
	Qty       uint32  `json:"qty"`
	Cost      float32 `json:"cost"`
}

type Order struct {
	OrderId     string      `json:"order_id"`
	CreatedDate string      `json:"create_date"`
	Cost        float32     `json:"cost"`
	Items       []OrderItem `json:"items"`
}
