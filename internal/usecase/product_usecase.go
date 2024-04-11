package usecase

type IProductUsecase interface {
	GetAllProductsByUserID() error
}
