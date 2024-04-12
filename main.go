package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/y-miyakaw/price-comparison-api/internal/controller"
	"github.com/y-miyakaw/price-comparison-api/internal/repository"
	"github.com/y-miyakaw/price-comparison-api/internal/router"
	"github.com/y-miyakaw/price-comparison-api/internal/usecase"
)

func main() {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}
	log.Println(db)
	productRepository := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productController := controller.NewProductController(productUsecase)
	e := router.NewRouter(productController)
	e.Logger.Fatal(e.Start(":8080"))
}
