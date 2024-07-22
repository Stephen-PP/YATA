package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stephen-pp/yata/api/internal/models"
)

func GetDB() *sql.DB {
	driver := os.Getenv("DB_DRIVER")
	connection := os.Getenv("DB_CONNECTION")

	switch driver {
	case "sqlite3":
		return initializeDB(driver, connection)
	case "mysql":
		return initializeDB(driver, connection)
	default:
		log.Fatal("unsupported driver")
		return nil
	}
}

func initializeDB(driver, connection string) *sql.DB {
	db, err := sql.Open(driver, connection)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTables(db, driver)

	return db
}

func createTables(db *sql.DB, driver string) {
	err := models.CreateAccessTokenTable(db, driver)
	if err != nil {
		log.Fatal(err)
	}
}
