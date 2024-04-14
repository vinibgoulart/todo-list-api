package task

type TaskStruct struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Priority    TaskPriority `json:"priority"`
	Status      TaskStatus   `json:"status"`
	ReleaseDate string       `json:"release_date"`
}
