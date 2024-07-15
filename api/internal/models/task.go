package models

import (
	"database/sql"
	"fmt"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ListID      int    `json:"list_id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    int    `json:"priority"`
	AssigneeID  int    `json:"assignee_id"`
	db          *sql.DB
}

func NewTask(db *sql.DB, name string, listID int, description string, status string, priority int, assigneeID int) *Task {
	return &Task{db: db, Name: name, ListID: listID, Description: description, Status: status, Priority: priority, AssigneeID: assigneeID}
}

func (t *Task) Save() error {
	if t.ID == 0 {
		return t.create(t.db)
	}
	return t.update(t.db)
}

func (t *Task) create(db *sql.DB) error {
	result, err := db.Exec("INSERT INTO tasks (name, list_id, description, status, priority, assignee_id) VALUES (?, ?, ?, ?, ?, ?)", t.Name, t.ListID, t.Description, t.Status, t.Priority, t.AssigneeID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	t.ID = int(id)
	return err
}

func (t *Task) update(db *sql.DB) error {
	_, err := db.Exec("UPDATE tasks SET name=?, list_id=?, description=?, status=?, priority=?, assignee_id=? WHERE id=?", t.Name, t.ListID, t.Description, t.Status, t.Priority, t.AssigneeID, t.ID)
	return err
}

func GetTaskByID(db *sql.DB) error {
	task := &Task{}
	row := db.QueryRow("SELECT id, name, list_id, description, status, priority, assignee_id FROM tasks WHERE id = ?", task.ID)
	return row.Scan(&task.ID, &task.Name, &task.ListID, &task.Description, &task.Status, &task.Priority, &task.AssigneeID)
}

func GetAllTasks(db *sql.DB, listID int, status string, priority string, assigneeID int) ([]Task, error) {
	query := "SELECT id, name, list_id, description, status, priority, assignee_id FROM tasks WHERE 1=1"
	if listID != 0 {
		query += " AND list_id = " + fmt.Sprint(listID)
	}
	if status != "" {
		query += " AND status = '" + status + "'"
	}
	if priority != "" {
		query += " AND priority = '" + priority + "'"
	}
	if assigneeID != 0 {
		query += " AND assignee_id = " + fmt.Sprint(assigneeID)
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []Task{}
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Name, &task.ListID, &task.Description, &task.Status, &task.Priority, &task.AssigneeID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
