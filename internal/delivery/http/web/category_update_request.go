package web

type CategoryUpdateRequest struct {
	Name string `json:"name" binding:"required,min=5,max=100"`
}
