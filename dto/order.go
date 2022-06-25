package dto

type OrderProductRequest struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Image    string  `json:"image"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
}

type OrderRequest struct {
	Name     string                `json:"name"`
	Tel      string                `json:"tel"`
	Products []OrderProductRequest `json:"products"`
}

type OrderProductResponse struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Image    string  `json:"image"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
}

type OrderResponse struct {
	ID       uint                   `json:"id"`
	Name     string                 `json:"name"`
	Tel      string                 `json:"tel"`
	Products []OrderProductResponse `json:"products"`
}
