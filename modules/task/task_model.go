package task

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type TaskStorage interface {
	Insert(*sql.DB, Task) (*Task, error)
}

type TaskStore struct {
}

type Task struct {
	ID          uuid.UUID    `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Priority    TaskPriority `json:"priority"`
	Status      TaskStatus   `json:"status"`
	ReleaseDate time.Time    `json:"release_date"`
}

func (t *TaskStore) Insert(db *sql.DB, task Task) (*Task, error) {
	sql := `INSERT INTO tasks (name, description, priority, release_date, status) VALUES ($1, $2, $3, $4, $5) RETURNING *`

	fmt.Println(task)
	row := db.QueryRow(sql, task.Name, task.Description, task.Priority, task.ReleaseDate, task.Status)

	taskCreated := &Task{}

	fmt.Println(taskCreated)

	err := row.Scan(&taskCreated.ID, &taskCreated.Name, &taskCreated.Description, &taskCreated.Priority, &taskCreated.ReleaseDate, &taskCreated.Status)
	if err != nil {
		return nil, err
	}

	return taskCreated, nil
}
