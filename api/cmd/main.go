package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stephen-pp/yata/api/db"
	"github.com/stephen-pp/yata/api/internal/models"
)

func main() {
	// SQLite
	dbSQLite := db.InitializeDB("sqlite3", "./app.db")
	defer dbSQLite.Close()

	// MySQL
	dbMySQL := db.InitializeDB("mysql", "root:password@tcp(localhost:3306)/test")
	defer dbMySQL.Close()

	// Create a new category
	category := models.NewCategory(dbMySQL, "Personal") // Replace dbSQLite with dbMySQL to test MySQL
	err := category.Save()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data created successfully!")
}
