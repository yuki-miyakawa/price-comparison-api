package router

import (
	"log"
	"os"

	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		log.Println(os.Getenv("TEST"))
		return c.String(200, "this is a healthcheck response from the server")
	})
	p := e.Group("/products")
	p.GET("/:userID", getAllProductsByUserID())

	return e
}
