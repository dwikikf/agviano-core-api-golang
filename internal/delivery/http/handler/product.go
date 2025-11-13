package handler

import (
	"net/http"
	"strconv"

	"github.com/dwikikf/agviano-core-api-golang/internal/delivery/http/web"
	domainProd "github.com/dwikikf/agviano-core-api-golang/internal/domain/product"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	uc domainProd.Usecase
}

func NewProductHandler(uc domainProd.Usecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) FindAll(c *gin.Context) {
	res, err := h.uc.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get products",
			Data:    nil,
		})
		return
	}

	var productsResp []*web.ProductResponse
	for _, prod := range res {
		productsResp = append(productsResp, web.ToProductResponse(prod))
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success get all products",
		Data:    productsResp,
	})
}

func (h *ProductHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	prodID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid product ID",
			Data:    nil,
		})
		return
	}

	prod, err := h.uc.GetByID(c.Request.Context(), prodID)
	if err != nil {
		if err == domainProd.ErrNotFound {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:    http.StatusNotFound,
				Message: "Product not found",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get product",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success get product by ID",
		Data:    web.ToProductResponse(*prod),
	})
}

func (h *ProductHandler) Create(c *gin.Context) {
	req := web.ProductCreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Data:    nil,
		})
		return
	}

	input := &domainProd.CreateProductInput{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Slug:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    "",
		IsActive:    true,
	}

	newProd, err := h.uc.Create(c.Request.Context(), input)
	if err != nil {
		// log.Println(err)
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create product",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, web.WebResponse{
		Code:    http.StatusCreated,
		Message: "Product created successfully",
		Data:    web.ToProductResponse(*newProd),
	})
}

func (h *ProductHandler) Update(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid product ID",
			Data:    nil,
		})
		return
	}

	var req web.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Data:    nil,
		})
		return
	}

	input := &domainProd.UpdateProductInput{
		ID:          idUint,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Slug:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    "",
		IsActive:    true,
	}

	updatedProd, err := h.uc.Update(c.Request.Context(), input)
	if err != nil {
		if err == domainProd.ErrNotFound {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:    http.StatusNotFound,
				Message: "Product not found",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update product",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Product updated successfully",
		Data:    web.ToProductResponse(*updatedProd),
	})
}

func (h *ProductHandler) Delete(c *gin.Context) {
	// Implementation goes here
}
