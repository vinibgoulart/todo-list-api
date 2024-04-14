package task

type TaskStruct struct {
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Priority    TaskPriority `json:"priority,omitempty"`
	Status      TaskStatus   `json:"status,omitempty"`
	ReleaseDate string       `json:"release_date,omitempty"`
}
