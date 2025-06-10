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

func (r *ProductRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	query := "SELECT COUNT(*) FROM products"
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	return count, err
}

func (r *ProductRepository) GetList(ctx context.Context, offset int, limit int) ([]*ProductModel, error) {
	query := "SELECT id, name, category, price FROM products ORDER BY name LIMIT ? OFFSET ?"

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*ProductModel, 0)
	for rows.Next() {
		var p ProductModel
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, rows.Err()
}

func (r *ProductRepository) GetById(ctx context.Context, id string) (*ProductModel, error) {
	query := "SELECT id, name, category, price FROM products WHERE id = ?"

	var p ProductModel
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Category, &p.Price)
	if err != nil {
		return nil, err
	}

	product := &p

	return product, err
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

func (r *ProductRepository) Update(ctx context.Context, p *ProductModel) (interface{}, error) {
	query := "UPDATE products SET name=?, category=?, price=? WHERE id=?"

	result, err := r.db.ExecContext(ctx, query, &p.Name, &p.Category, &p.Price, &p.ID)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *ProductRepository) Delete(ctx context.Context, p *ProductModel) (interface{}, error) {
	query := "DELETE FROM products WHERE id=?;"

	result, err := r.db.ExecContext(ctx, query, p.ID)

	return result, err
}