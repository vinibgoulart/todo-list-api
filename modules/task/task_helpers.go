package task

import (
	"errors"
	"time"
)

func TaskApiToModelMapping(taskStruct *TaskStruct) (*Task, error) {
	releaseDate, err := time.Parse(time.RFC3339, taskStruct.ReleaseDate)
	if err != nil {
		return nil, err
	}

	if !TaskPriorityValid[taskStruct.Priority] {
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
