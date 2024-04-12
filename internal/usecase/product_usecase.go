package usecase

import (
	"github.com/y-miyakaw/price-comparison-api/gen/sqlc"
	"github.com/y-miyakaw/price-comparison-api/internal/repository"
)

type IProductUsecase interface {
	GetAllProductsByUserID(userID string) ([]sqlc.Product, error)
}

type productUsecase struct {
	pr repository.IProductRepository
}

func NewProductUsecase(pr repository.IProductRepository) IProductUsecase {
	return &productUsecase{pr}
}

func (pu *productUsecase) GetAllProductsByUserID(userID string) ([]sqlc.Product, error) {
	products, err := pu.pr.GetAllProductsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
