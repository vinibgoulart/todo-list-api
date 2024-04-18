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

	var taskStatus TaskStatus

	if taskStruct.Status == "" {
		taskStatus = TaskStatus(TaskStatusOpen)
	} else {
		taskStatus = taskStruct.Status
	}

	return &Task{
		Name:        taskStruct.Name,
		Description: taskStruct.Description,
		Priority:    TaskPriority(taskStruct.Priority),
		Status:      taskStatus,
		ReleaseDate: releaseDate,
	}, nil
}
