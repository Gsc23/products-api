package handler

import (
	"context"
	"example/products/internal/product/command"
	"example/products/internal/product/repository"
	"fmt"
)

type ProductHandler struct {
	repository *repository.ProductRepository
}

func NewProductHandler(repository *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repository: repository}
}

func (p *ProductHandler) Handle(ctx context.Context, msg interface{}) (interface{}, error) {
	switch cmd := msg.(type) {
	case command.CreateProductCommand:
		fmt.Printf("Handler recebendo comando para criar produto: %s\n", cmd.Name)

		newProduct := &repository.ProductModel{
			Name:     cmd.Name,
			Category: cmd.Category,
			Price:    cmd.Price,
		}

		product, err := p.repository.Save(ctx, newProduct)
		if err != nil {
			return nil, err
		}

		return product.ID, nil // Retorna o ID do produto criado
	default:
		return nil, fmt.Errorf("ProductHandler não sabe como lidar com o tipo: %T", msg)

	}
}
