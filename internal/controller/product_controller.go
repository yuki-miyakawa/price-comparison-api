package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/y-miyakaw/price-comparison-api/internal/usecase"
)

type IProductController interface {
	GetAllProductsByUserID(c echo.Context) error
}

type productController struct {
	pu usecase.IProductUsecase
}

func NewProductController(pu usecase.IProductUsecase) IProductController {
	return &productController{pu}
}

func (pc productController) GetAllProductsByUserID(c echo.Context) error {
	userID := c.Param("userID")
	products, err := pc.pu.GetAllProductsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, products)
}
