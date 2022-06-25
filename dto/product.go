package dto

type ProductRequest struct {
	SKU        string  `form:"sku"`
	Name       string  `form:"name"`
	Desc       string  `form:"desc"`
	Price      float64 `form:"price"`
	Status     uint    `form:"status"`
	CategoryID uint    `form:"categoryId"`
}

type CreateOrUpdateProductResponse struct {
	ID         uint    `json:"id"`
	SKU        string  `json:"sku"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float64 `json:"price"`
	Status     uint    `json:"status"`
	CategoryID uint    `json:"categoryId"`
	Image      string  `json:"image"`
}

type ReadProductResponse struct {
	ID       uint             `json:"id"`
	SKU      string           `json:"sku"`
	Name     string           `json:"name"`
	Desc     string           `json:"desc"`
	Price    float64          `json:"price"`
	Status   uint             `json:"status"`
	Image    string           `json:"image"`
	Category CategoryResponse `json:"category"`
}
