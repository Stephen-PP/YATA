package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stephen-pp/yata/api/internal/handlers"
)

func main() {
	// Create a new category
	// category := models.NewCategory(dbSQLite, "Personal") // Replace dbSQLite with dbMySQL to test MySQL
	// err := category.Save()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("starting webserver on port :8080")

	// Register functions and middlewares
	http.HandleFunc("/access-token", handlers.GenerateAccessToken)

	err := http.ListenAndServe(":8080", nil)
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
