package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/y-miyakaw/price-comparison-api/gen/sqlc"
)

type IProductRepository interface {
	GetAllProductsByUserID(userID string) ([]sqlc.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) IProductRepository {
	return &productRepository{db}
}

func (pr *productRepository) GetAllProductsByUserID(userID string) ([]sqlc.Product, error) {
	q := sqlc.New(pr.db)
	ctx := context.Background()
	products, err := q.GetAllProductsByUserID(ctx, userID)
	if err != nil {
		log.Printf("error GetAllProductsByUserID: %v", err)
		return nil, err
	}
	return products, nil
}
