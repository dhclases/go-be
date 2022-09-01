package handler

import (
	"actividad/clase/internal/domain"
	"actividad/clase/internal/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	s product.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewProductHandler(s product.Service) *productHandler {
	return &productHandler{
		s: s,
	}
}

func InitProductHandler() *productHandler {
	// Inicializa producto
	list := product.LoadProducts("../products.json")
	repo := product.NewRepository(list)
	service := product.NewService(repo)
	return NewProductHandler(service)
}

// GetAll obtiene todos los productos
func (h *productHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, _ := h.s.GetAll()
		c.JSON(200, products)
	}
}

// GetByID obtiene un producto por su id
func (h *productHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		product, err := h.s.GetByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "product not found"})
			return
		}
		c.JSON(200, product)
	}
}

// Search busca un producto por precio mayor a un valor
func (h *productHandler) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		priceParam := c.Query("priceGt")
		price, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid price"})
			return
		}
		products, err := h.s.SearchPriceGt(price)
		if err != nil {
			c.JSON(404, gin.H{"error": "no products found"})
			return
		}
		c.JSON(200, products)
	}
}

// Post crear un producto nuevo
func (h *productHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newProduct domain.Product
		err := ctx.ShouldBindJSON(&newProduct)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid product"})
			return
		}
		valid, err := product.ValidateEmptys(&newProduct)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		valid, err = product.ValidateExpiration(&newProduct)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		p, err := h.s.Create(newProduct)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)
	}
}

// Post crear un producto nuevo
func (h *productHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newProduct domain.Product
		err := ctx.ShouldBindJSON(&newProduct)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid product"})
			return
		}
		valid, err := product.ValidateEmptys(&newProduct)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		valid, err = product.ValidateExpiration(&newProduct)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		p, err := h.s.Update(newProduct)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)
	}
}

// DeleteById elimina un producto por su id
func (h *productHandler) DeleteByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = h.s.DeleteByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "product not found"})
			return
		}
		c.JSON(204, nil)
	}
}
