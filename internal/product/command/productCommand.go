package command

import "example/products/internal/product/repository"

type ListProductsQuery struct {
	Page     int
	PageSize int
}

type PaginatedProductsResult struct {
	Products     []*repository.ProductModel `json:"products"`
	Page         int                 `json:"page"`
	PageSize     int                 `json:"pageSize"`
	TotalRecords int64               `json:"totalRecords"`
	TotalPages   int                 `json:"totalPages"`
}

type CreateProductCommand struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int32  `json:"price"`
}