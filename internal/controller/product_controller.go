package controller

import "github.com/labstack/echo"

type IProductController interface {
	GetAllProductsByUserID(c echo.Context) error
}

type ProductController struct {
}

func GetAllProductsByUserID(c echo.Context) error {
	return c.String(200, "this is a response from the server")
}
