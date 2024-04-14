package task

import (
	"time"

	"github.com/gofrs/uuid"
)

type TaskStorage interface {
}

type TaskStore struct {
}

type Task struct {
	ID          uuid.UUID `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	ReleaseDate time.Time `json:"release_date"`
}
