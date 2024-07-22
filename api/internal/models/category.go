package models

import (
	"database/sql"
	"errors"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	db   *sql.DB
}

func NewCategory(db *sql.DB, name string) *Category {
	return &Category{Name: name, db: db}
}

func (c *Category) Save() error {
	if c.ID == 0 {
		return c.create()
	}
	return c.update()
}

func (c *Category) create() error {
	result, err := c.db.Exec("INSERT INTO categories (name) VALUES (?)", c.Name)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	c.ID = int(id)
	return err
}

func (c *Category) update() error {
	_, err := c.db.Exec("UPDATE categories SET name=? WHERE id=?", c.Name, c.ID)
	return err
}

func GetCategoryByID(db *sql.DB, id int) (*Category, error) {
	category := &Category{db: db}
	row := db.QueryRow("SELECT id, name FROM categories WHERE id = ?", id)
	err := row.Scan(&category.ID, &category.Name)
	if err == sql.ErrNoRows {
		return nil, errors.New("category not found")
	} else if err != nil {
		return nil, err
	}
	return category, nil
}

func GetAllCategories(db *sql.DB) ([]*Category, error) {
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []*Category{}
	for rows.Next() {
		category := &Category{db: db}
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func CreateCategoryTable(db *sql.DB, db_type string) error {
	if db_type == "sqlite3" {
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
		);`)
		return err
	} else if db_type == "mysql" {
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY AUTO_INCREMENT,
			name TEXT NOT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
		return err
	} else {
		return errors.New("unsupported driver")
	}
}
