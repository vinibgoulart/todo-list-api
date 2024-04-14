package task

type TaskStatus string

const (
	TaskStatusCompleted  TaskStatus = "COMPLETED"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusOpen       TaskStatus = "OPEN"
)

var TaskStatusValid = map[TaskStatus]bool{
	TaskStatusCompleted:  true,
	TaskStatusInProgress: true,
	TaskStatusOpen:       true,
}
