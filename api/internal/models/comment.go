package models

import (
	"database/sql"
)

type Comment struct {
	ID     int    `json:"id"`
	TaskID int    `json:"task_id"`
	Text   string `json:"text"`
	db     *sql.DB
}

func NewComment(db *sql.DB, taskID int, text string) *Comment {
	return &Comment{TaskID: taskID, Text: text, db: db}
}

func (c *Comment) Save() error {
	if c.ID == 0 {
		return c.create()
	} else {
		return c.update()
	}
}

func (c *Comment) create() error {
	result, err := c.db.Exec("INSERT INTO comments (task_id, text) VALUES (?, ?)", c.TaskID, c.Text)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	c.ID = int(id)
	return err
}

func (c *Comment) update() error {
	_, err := c.db.Exec("UPDATE comments SET task_id=?, text=? WHERE id=?", c.TaskID, c.Text, c.ID)
	return err
}

func GetCommentByID(db *sql.DB, id int) (*Comment, error) {
	comment := &Comment{db: db}
	row := db.QueryRow("SELECT id, task_id, text FROM comments WHERE id = ?", id)
	return comment, row.Scan(&comment.ID, &comment.TaskID, &comment.Text)
}

func GetAllCommentsByTaskID(db *sql.DB, taskID int) ([]Comment, error) {
	rows, err := db.Query("SELECT id, task_id, text FROM comments WHERE task_id = ?", taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments := []Comment{}
	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.ID, &comment.TaskID, &comment.Text)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
