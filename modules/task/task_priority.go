package task

type TaskPriority string

const (
	TaskPriorityHigh   TaskPriority = "HIGH"
	TaskPriorityMedium TaskPriority = "MEDIUM"
	TaskPriorityLow    TaskPriority = "LOW"
)

var TaskPriorityValid = map[TaskPriority]bool{
	TaskPriorityHigh:   true,
	TaskPriorityMedium: true,
	TaskPriorityLow:    true,
}
