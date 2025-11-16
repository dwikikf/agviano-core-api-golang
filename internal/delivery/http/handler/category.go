package handler

import (
	"net/http"
	"strconv"

	"github.com/dwikikf/agviano-core-api-golang/internal/delivery/http/web"
	domainCat "github.com/dwikikf/agviano-core-api-golang/internal/domain/category"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	uc domainCat.Usecase
}

func NewCategoryHandler(uc domainCat.Usecase) *CategoryHandler {
	return &CategoryHandler{uc}
}

func (h *CategoryHandler) FindAll(c *gin.Context) {
	list, err := h.uc.GetAll(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	var res []*web.CategoryResponse
	for _, cat := range list {
		res = append(res, web.ToCategoryResponse(cat))
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success get all categories",
		Data:    res,
	})
}

func (h *CategoryHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid category ID",
		})
		return
	}

	cat, err := h.uc.GetByID(c.Request.Context(), idUint)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success get category by ID",
		Data:    web.ToCategoryResponse(*cat),
	})
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req web.CategoryCreateRequest

	// validation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err) //midlleware
		return
	}

	// logic
	newCat, err := h.uc.Create(
		c.Request.Context(),
		&domainCat.CreateCatData{Name: req.Name})

	if err != nil {
		c.Error(err) //midlleware
		return
	}

	// response
	c.JSON(http.StatusCreated, web.WebResponse{
		Code:    http.StatusCreated,
		Message: "Category created successfully",
		Data:    web.ToCategoryResponse(*newCat),
	})
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid category ID",
		})
		return
	}

	var req web.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	updatedCat, err := h.uc.Update(c.Request.Context(), &domainCat.UpdateCatData{
		ID:   idUint,
		Name: req.Name,
	})

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Category updated successfully",
		Data:    web.ToCategoryResponse(*updatedCat),
	})
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	// implementation for deleting a category
}
