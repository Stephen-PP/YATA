package models

import (
	"database/sql"
	"errors"
)

type List struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description"`
	db          *sql.DB
}

func NewList(db *sql.DB, name string, categoryID int, description string) *List {
	return &List{Name: name, CategoryID: categoryID, Description: description, db: db}
}

func (l *List) Save() error {
	if l.ID == 0 {
		return l.create(l.db)
	}
	return l.update(l.db)
}

func (l *List) create(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO lists (name, category_id, description) VALUES (?, ?, ?)", l.Name, l.CategoryID, l.Description)
	return err
}

func (l *List) update(db *sql.DB) error {
	_, err := db.Exec("UPDATE lists SET name=?, category_id=?, description=? WHERE id=?", l.Name, l.CategoryID, l.Description, l.ID)
	return err
}

func GetListByID(db *sql.DB) error {
	list := &List{db: db}
	row := db.QueryRow("SELECT id, name, category_id, description FROM lists WHERE id = ?", list.ID)
	return row.Scan(&list.ID, &list.Name, &list.CategoryID, &list.Description)
}

func GetAllListsByCategory(db *sql.DB, categoryID int) ([]List, error) {
	rows, err := db.Query("SELECT id, name, category_id, description FROM lists WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lists := []List{}
	for rows.Next() {
		var list List
		err = rows.Scan(&list.ID, &list.Name, &list.CategoryID, &list.Description)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

func CreateListTable(db *sql.DB, db_type string) error {
	if db_type == "sqlite3" {
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS lists (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			category_id INTEGER NOT NULL,
			description TEXT NOT NULL
		);`)
		return err
	} else if db_type == "mysql" {
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS lists (
			id INTEGER PRIMARY KEY AUTO_INCREMENT,
			name TEXT NOT NULL,
			category_id INTEGER NOT NULL,
			description TEXT NOT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
		return err
	} else {
		return errors.New("unsupported driver")
	}
}
