package handler

import (
	"context"
	"example/products/internal/product/command"
	"example/products/internal/product/repository"
	"fmt"
	"log"
	"math"
)

type ProductHandler struct {
	repository *repository.ProductRepository
}

func NewProductHandler(repository *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repository: repository}
}

func (p *ProductHandler) Handle(ctx context.Context, msg interface{}) (interface{}, error) {
	switch cmd := msg.(type) {
		case command.ListProductsQuery:
			log.Printf("Handler recebendo query para listar produtos: Página %d\n", cmd.Page)

			totalRecords, err := p.repository.Count(ctx)
			if err != nil {
				return nil, err
			}

			offset := (cmd.Page - 1) * cmd.PageSize
			products, err := p.repository.GetList(ctx, offset, cmd.PageSize)
			if err != nil {
				return nil, err
			}

			return &command.PaginatedProductsResult{
				Products:     products,
				Page:         cmd.Page,
				PageSize:     cmd.PageSize,
				TotalRecords: totalRecords,
				TotalPages:   int(math.Ceil(float64(totalRecords) / float64(cmd.PageSize))),
			}, nil
			
		case command.CreateProductCommand:
			log.Printf("Handler recebendo comando para criar produto: %s\n", cmd.Name)

			newProduct := &repository.ProductModel{
				Name:     cmd.Name,
				Category: cmd.Category,
				Price:    cmd.Price,
			}

			product, err := p.repository.Save(ctx, newProduct)
			if err != nil {
				return nil, err
			}

			return product.ID, nil
		default:
			return nil, fmt.Errorf("ProductHandler não sabe como lidar com o tipo: %T", msg)
	}
}
