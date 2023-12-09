package main

import (
	"fmt"

	"github.com/y-miyakaw/price-comparison-api/cmd/db"
	"github.com/y-miyakaw/price-comparison-api/cmd/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Migration finished")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.Product{})
}
