package task

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type TaskStorage interface {
	Insert(*sql.DB, Task) (*Task, error)
	GetAll(*sql.DB) ([]*Task, error)
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
