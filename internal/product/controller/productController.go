package controller

import (
	"example/products/internal/product/command"
	"example/products/pkg/bus"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	bus *bus.Bus
}

func NewProductController(b *bus.Bus) *productController {
	return &productController{bus: b}
}

func (pc *productController) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	qry := command.ListProductsQuery{
		Page:     page,
		PageSize: pageSize,
	}

	result, err := pc.bus.Dispatch(c.Request.Context(), qry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (pc *productController) GetProduct(c *gin.Context) {
	productId, _ := c.Params.Get("id")

	cmd := command.GetProductCommand{
		ID: productId,
	}

	result, err := pc.bus.Dispatch(c.Request.Context(), cmd)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
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

func (pc *productController) Update(c *gin.Context) {
	productId, _ := c.Params.Get("id")

	cmd := command.UpdateProductCommand{
		ID: productId,
	}
	
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := pc.bus.Dispatch(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Produto atualizados",
		"produto": result,
	})
}

func (pc *productController) Delete(c *gin.Context) {
	productId, _ := c.Params.Get("id")

	cmd := command.UpdateProductCommand{
		ID: productId,
	}

	_, err := pc.bus.Dispatch(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Produto deletado",
	})
}