package task

import (
	"errors"
	"fmt"
	taskDomain "tracker_backend/src/domain/task"
	"tracker_backend/src/domain/user"
)

var (
	ErrCreateTask = errors.New("failed to create task")
)

type TaskInCreate struct {
	Description   string
	OwnerUsername string
}

type CreatedTaskArtefacts struct {
	TaskNumber int
	Stage      string
}

type CreateTaskCmd struct {
	Saver TaskSaver
}

func (c CreateTaskCmd) Execute(taskDto TaskInCreate) (CreatedTaskArtefacts, error) {
	description := taskDomain.Description(taskDto.Description)
	err := description.Validate()
	if err != nil {
		return CreatedTaskArtefacts{}, err
	}
	ownerUsername := user.Username(taskDto.OwnerUsername)
	err = ownerUsername.Validate()
	if err != nil {
		return CreatedTaskArtefacts{}, err
	}
	stage := string(taskDomain.ToDo)
	taskNumber, err := c.Saver.SaveIncrTaskNumber(TaskSaveDto{
		Description:   taskDto.Description,
		OwnerUsername: taskDto.OwnerUsername,
		Stage:         stage,
	})
	if err != nil {
		return CreatedTaskArtefacts{}, fmt.Errorf("%w: %s", ErrCreateTask, err)
	}
	return CreatedTaskArtefacts{
		TaskNumber: taskNumber,
		Stage:      stage,
	}, nil
}
