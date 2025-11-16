package web

type ProductUpdateRequest struct {
	Name        string  `json:"name" binding:"required,min=10,max=150"`
	Description string  `json:"description" binding:"required,min=10"`
	Price       float64 `json:"price" binding:"required"`
	Stock       uint    `json:"stock" binding:"required"`
	CategoryID  uint64  `json:"category_id" binding:"required"`
}
