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
		log.Printf("ProductHandler receiving command for list products: Página %d\n", cmd.Page)

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
	case command.GetProductCommand:
		log.Printf("ProductHandler receiving command for list products: ¨%s\n", cmd.Name)

		product, err := p.repository.GetById(ctx, cmd.ID)
		if err != nil {
			return nil, err
		}

		return &command.GetProductCommand{
			ID:       product.ID,
			Name:     product.Name,
			Category: product.Category,
			Price:    product.Price,
		}, nil

	case command.CreateProductCommand:
		log.Printf("ProductHandler receiving command for create products: %s\n", cmd.Name)

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
	case command.UpdateProductCommand:
		log.Printf("ProductHandler receiving command for update products with ID: %s\n", cmd.ID)

		product := &repository.ProductModel{
			ID:       cmd.ID,
			Name:     cmd.Name,
			Category: cmd.Category,
			Price:    cmd.Price,
		}

		result, err := p.repository.Update(ctx, product)
		if err != nil {
			return nil, err
		}

		return result, nil
	case command.DeleteProductCommand:
		log.Printf("ProductHandler receiving command for update products with ID: %s\n", cmd.ID)

		product := &repository.ProductModel{
			ID:       cmd.ID,
		}

		result, err := p.repository.Delete(ctx, product)
		if err != nil {
			return nil, err
		}

		return result, nil
	default:
		return nil, fmt.Errorf("ProductHandler doesn't know what to do with: %T", msg)
	}
}
