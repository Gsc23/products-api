package controller

import (
	"example/products/internal/product/command"
	"example/products/pkg/bus"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	bus *bus.Bus
}

func NewProductController(b *bus.Bus) *productController {
	return &productController{bus: b}
}

func (pc *productController) Create(c *gin.Context) {

	var cmd command.CreateProductCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productId, err := pc.bus.Dispatch(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Produto Criado",
		"id":      productId,
	})
}
