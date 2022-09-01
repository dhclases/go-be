package handler

import "github.com/gin-gonic/gin"

func NewProductRouter(e *gin.Engine, h productHandler) {

	products := e.Group("/api/v1/products")
	{
		products.GET("", h.GetAll())
		products.GET(":id", h.GetByID())
		products.GET("/search", h.Search())
		products.POST("", h.Create())
		products.PUT("", h.Update())
		products.DELETE(":id", h.DeleteByID())
	}
}
