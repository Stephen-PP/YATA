package models

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	db       *sql.DB
}

func NewUser(db *sql.DB, username, email, password string) *User {
	return &User{Username: username, Email: email, Password: password, db: db}
}

func (u *User) Save() error {
	if u.ID == 0 {
		return u.create()
	}
	return u.update()
}

func (u *User) create() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	result, err := u.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", u.Username, u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = int(id)
	return err
}

func (u *User) update() error {
	_, err := u.db.Exec("UPDATE users SET username=?, email=?, password=? WHERE id=?", u.Username, u.Email, u.Password, u.ID)
	return err
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	user := &User{db: db}
	row := db.QueryRow("SELECT id, username, email, password FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	user := &User{db: db}
	row := db.QueryRow("SELECT id, username, email, password FROM users WHERE username = ?", username)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	return user, nil
}
