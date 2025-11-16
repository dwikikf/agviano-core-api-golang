package middleware

import (
	"errors"
	"net/http"

	"github.com/dwikikf/agviano-core-api-golang/internal/delivery/http/web"
	"github.com/dwikikf/agviano-core-api-golang/internal/errs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		e := err.Err

		// Gin Validation Error
		var valErr validator.ValidationErrors
		if errors.As(e, &valErr) {
			c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:    http.StatusBadRequest,
				Message: "validation error",
				Error:   web.FormatValidationError(e),
			})
			return
		}

		// Domain Errors
		switch {
		case errors.Is(e, errs.ErrNotFound):
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:    http.StatusNotFound,
				Message: "record not found",
			})
			return

		case errors.Is(e, errs.ErrConflict):
			c.JSON(http.StatusConflict, web.WebResponse{
				Code:    http.StatusConflict,
				Message: "duplicate entry data",
			})
			return

			// case errors.Is(e, errs.ErrInvalid):
			// 	c.JSON(http.StatusBadRequest, web.WebResponse{
			// 		Code:    http.StatusBadRequest,
			// 		Message: "invalid data",
			// 	})
			// 	return
		}

		// Default
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Error:   e.Error(),
		})
	}
}
