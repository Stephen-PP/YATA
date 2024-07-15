package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func InitializeDB(driver, connection string) *sql.DB {
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
	var query string
	switch driver {
	case "sqlite3":
		query = `
            CREATE TABLE IF NOT EXISTS categories (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL
            );
        `
	case "mysql":
		query = `
            CREATE TABLE IF NOT EXISTS categories (
                id INT AUTO_INCREMENT,
                name VARCHAR(255) NOT NULL,
                PRIMARY KEY (id)
            ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
        `
	default:
		log.Fatal("Unsupported driver")
	}
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
