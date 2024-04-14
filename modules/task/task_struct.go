package task

type TaskStruct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	ReleaseDate string `json:"release_date"`
}
