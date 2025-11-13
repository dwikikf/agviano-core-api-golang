package web

type CategoryUpdateRequest struct {
	Name string `json:"name" binding:"required" validate:"min=1,max=100"`
}
