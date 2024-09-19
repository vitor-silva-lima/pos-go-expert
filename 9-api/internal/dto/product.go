package dto

type ProductDtoInputBase struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type CreateProductDtoInput struct {
	ProductDtoInputBase
}

type UpdateProductDtoInput struct {
	ProductDtoInputBase
	ProductID string `json:"product_id"`
}
