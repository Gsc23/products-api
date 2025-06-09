package main

import (
	"example/products/internal/db"
	"example/products/internal/router"
)

func main() {
	db := db.InitializeDatabase()
	defer db.Close()
	
	router.ProductRoutes(db)
}