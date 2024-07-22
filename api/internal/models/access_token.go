package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"log"
	"time"
)

type AccessToken struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	db        *sql.DB
}

func NewAccessToken(db *sql.DB, username string, id string) *AccessToken {
	return &AccessToken{Username: username, ID: id, db: db}
}

func (at *AccessToken) Save() error {
	at.generateToken()
	_, err := at.db.Exec("INSERT INTO access_tokens (id, username, token, created_at) VALUES (?, ?, ?, ?)", at.ID, at.Username, at.Token, time.Now())
	return err
}

func (at *AccessToken) generateToken() {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	at.Token = base64.StdEncoding.EncodeToString(b)
}

func CreateAccessTokenTable(db *sql.DB, db_type string) error {
	if db_type == "sqlite3" {
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS access_tokens (
				id TEXT NOT NULL,
				username TEXT NOT NULL,
				token TEXT PRIMARY KEY,
				created_at DATETIME NOT NULL
			);`)
		return err
	} else if db_type == "mysql" {
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS access_tokens (
				id VARCHAR(255) NOT NULL,
				username VARCHAR(255) NOT NULL,
				token VARCHAR(255) PRIMARY KEY,
				created_at DATETIME NOT NULL
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
		return err
	} else {
		return errors.New("unsupported driver")
	}
}
