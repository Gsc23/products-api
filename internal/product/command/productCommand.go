package command

import "example/products/internal/product/repository"

type ListProductsQuery struct {
	Page     int
	PageSize int
}

type PaginatedProductsResult struct {
	Products     []*repository.ProductModel `json:"products"`
	Page         int                        `json:"page"`
	PageSize     int                        `json:"pageSize"`
	TotalRecords int64                      `json:"totalRecords"`
	TotalPages   int                        `json:"totalPages"`
}

type GetProductCommand struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int32  `json:"price"`
}


type CreateProductCommand struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int32  `json:"price"`
}

type UpdateProductCommand struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int32  `json:"price"`
}

type DeleteProductCommand struct {
	ID       string `json:"id"`
}