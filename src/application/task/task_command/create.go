package task_command

import (
	"errors"
	"fmt"
	"tracker_backend/src/domain"
)

var (
	ErrCreateTask = errors.New("failed to create task")
)

type TaskInCreate struct {
	Description   string
	OwnerUsername string
}

type CreatedTaskArtefacts struct {
	TaskId domain.TaskId
	Stage  domain.Stage
}

type CreateTaskCmd struct {
	Saver TaskSaver
}

func (c CreateTaskCmd) Execute(taskDto TaskInCreate) (CreatedTaskArtefacts, error) {
	task, err := domain.NewTask(taskDto.Description, taskDto.OwnerUsername)
	if err != nil {
		return CreatedTaskArtefacts{}, err
	}
	newId, err := c.Saver.SaveIncrId(task)
	if err != nil {
		return CreatedTaskArtefacts{}, fmt.Errorf("%w: %s", ErrCreateTask, err)
	}
	err = task.SetConfirmedNewId(newId)
	if err != nil {
		return CreatedTaskArtefacts{}, err
	}
	id, ok := task.GetConfirmedId()
	if !ok {
		return CreatedTaskArtefacts{}, fmt.Errorf("%w: task id is not confirmed", ErrCreateTask)
	}
	return CreatedTaskArtefacts{
		TaskId: id,
		Stage:  task.GetStage(),
	}, nil
}
