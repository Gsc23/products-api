package router

import (
	"database/sql"
	"example/products/internal/product/command"
	"example/products/internal/product/controller"
	"example/products/internal/product/handler"
	"example/products/internal/product/repository"
	"example/products/pkg/bus"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(db *sql.DB) {
	productRepo := repository.NewProductRepository(db)

	productHandler := handler.NewProductHandler(productRepo)

	commandBus := bus.NewBus()
	commandBus.RegisterHandler(productHandler, command.CreateProductCommand{})
	commandBus.RegisterHandler(productHandler, command.ListProductsQuery{})

	productController := controller.NewProductController(commandBus)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/products", productController.Index)

	router.GET("/products/:id", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "product",
		})
	})

	router.POST("/products", productController.Create)

	router.PUT("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "updated",
		})
	})

	router.DELETE("/products", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{})
	})

	router.Run()
}
