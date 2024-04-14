package task

import (
	"errors"
	"time"
)

func HelperTaskApiToModelMapping(taskStruct *TaskStruct) (*Task, error) {
	releaseDate, err := time.Parse(time.RFC3339, taskStruct.ReleaseDate)
	if err != nil && taskStruct.ReleaseDate != "" {
		return nil, err
	}

	if !TaskPriorityValid[taskStruct.Priority] && taskStruct.Priority != "" {
		return nil, errors.New("invalid priority value, use: HIGH, MEDIUM, LOW")
	}

	return &Task{
		Name:        taskStruct.Name,
		Description: taskStruct.Description,
		Priority:    TaskPriority(taskStruct.Priority),
		Status:      TaskStatus(TaskStatusOpen),
		ReleaseDate: releaseDate,
	}, nil
}
