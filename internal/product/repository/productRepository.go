package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type ProductModel struct {
	ID       string
	Name     string
	Category string
	Price    int32
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Migrate() error {
	createProductTable := `
	CREATE TABLE IF NOT EXISTS products (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		category TEXT NOT NULL,
		price INTEGER NOT NULL
	);`

	_, err := r.db.Exec(createProductTable)
	return err
}

func (r *ProductRepository) Save(ctx context.Context, p *ProductModel) (*ProductModel, error) {
	p.ID = uuid.New().String()
	query := "INSERT INTO products (id, name, category, price) VALUES (?, ?, ?, ?)"

	_, err := r.db.ExecContext(ctx, query, p.ID, p.Name, p.Category, p.Price)
	if err != nil {
		return nil, err 
	}

	return p, nil
}