package task

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	dbHelper "github.com/vinibgoulart/todo-list/db"
)

type TaskStorage interface {
	Insert(*sql.DB, Task) (*Task, error)
	GetAll(*sql.DB) ([]*Task, error)
	Update(*sql.DB, Task, string) (*Task, error)
	Remove(*sql.DB, string)
}

type TaskStore struct {
}

type Task struct {
	ID          uuid.UUID    `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Priority    TaskPriority `json:"priority,omitempty"`
	Status      TaskStatus   `json:"status,omitempty"`
	ReleaseDate time.Time    `json:"release_date,omitempty"`
}

func (t *TaskStore) Insert(db *sql.DB, task Task) (*Task, error) {
	sql := `INSERT INTO tasks (name, description, priority, release_date, status) VALUES ($1, $2, $3, $4, $5) RETURNING *`

	fmt.Println(task)
	row := db.QueryRow(sql, task.Name, task.Description, task.Priority, task.ReleaseDate, task.Status)

	taskCreated := &Task{}

	err := row.Scan(&taskCreated.ID, &taskCreated.Name, &taskCreated.Description, &taskCreated.Priority, &taskCreated.ReleaseDate, &taskCreated.Status)
	if err != nil {
		return nil, err
	}

	return taskCreated, nil
}

func (t *TaskStore) GetAll(db *sql.DB) ([]*Task, error) {
	sql := `SELECT * FROM tasks`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*Task{}

	for rows.Next() {
		task := &Task{}
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Priority, &task.ReleaseDate, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskStore) Update(db *sql.DB, task Task, id string) (*Task, error) {
	query, args := dbHelper.HelperDBUpdateTableById("tasks", task, id)

	row := db.QueryRow(query, args...)

	taskUpdated := &Task{}

	err := row.Scan(&taskUpdated.ID, &taskUpdated.Name, &taskUpdated.Description, &taskUpdated.Priority, &taskUpdated.ReleaseDate, &taskUpdated.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	return taskUpdated, nil
}

func (t *TaskStore) Remove(db *sql.DB, id string) {
	sql := `DELETE FROM tasks WHERE id=$1`

	db.QueryRow(sql, id)
}
