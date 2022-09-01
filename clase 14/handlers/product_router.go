package handlers

import (
	"actividad/clase/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewProductRouter(e *gin.Engine, list []domain.Product) {
	products := e.Group("/api/v1/products")
	{
		products.GET("", GetAllProducts(list))
		products.GET(":id", GetProduct(list))
		products.GET("/search", SearchProduct(list))
	}
}

func GetAllProducts(list []domain.Product) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, list)
	}
}

func GetProduct(list []domain.Product) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Capturar el path Param
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
		}

		// Simulacion de la capa de servicio

		for _, product := range list {
			if product.Id == id {
				ctx.JSON(200, product)
				return
			}

		}

		ctx.JSON(404, gin.H{"error": "product not found"})

	}

}

func SearchProduct(list []domain.Product) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Capturar Query params
		query := ctx.Query("priceGt")
		priceGt, err := strconv.ParseFloat(query, 32)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid price"})
		}

		// Simulacion de la capa de servicio
		result := []domain.Product{}

		for _, product := range list {
			if product.Price > priceGt {
				result = append(result, product)
			}
		}

		ctx.JSON(200, result)

	}
}
