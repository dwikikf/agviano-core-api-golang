package http

import (
	"github.com/dwikikf/agviano-core-api-golang/internal/delivery/http/handler"
	"github.com/dwikikf/agviano-core-api-golang/internal/delivery/http/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(CategoryHandler *handler.CategoryHandler, ProductHandler *handler.ProductHandler) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.Use(middleware.ErrorHandler())
	// r.Use(gin.Logger())

	category := r.Group("/categories")
	{
		category.GET("", CategoryHandler.FindAll)
		category.GET("/:id", CategoryHandler.FindByID)
		category.POST("/", CategoryHandler.Create)
		category.PUT("/:id", CategoryHandler.Update)
		// category.DELETE("/:id", CategoryHandler.Delete)
	}

	product := r.Group("/products")
	{
		product.GET("", ProductHandler.FindAll)
		product.GET("/:id", ProductHandler.FindByID)
		product.POST("/", ProductHandler.Create)
		product.PUT("/:id", ProductHandler.Update)
		product.DELETE("/", ProductHandler.Delete)
	}

	// ping check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
