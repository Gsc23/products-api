package main

import (
	"example/products/internal/db"
	"example/products/internal/product/repository"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	log.Println("Iniciando a migração do banco de dados...")
	database := db.InitializeDatabase()
	defer database.Close()

	productRepo := repository.NewProductRepository(database)
	if err := productRepo.Migrate(); err != nil {
		log.Fatalf("Falha ao executar a migração: %v", err)
	}

	log.Println("Migração concluída com sucesso!")
}
