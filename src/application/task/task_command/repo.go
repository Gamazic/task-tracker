package task_command

import (
	"errors"
	"tracker_backend/src/domain"
)

type TaskSaver interface {
	SaveIncrId(task domain.Task) (domain.TaskId, error)
}

var ErrTaskNotFound = errors.New("task not found")

type TaskStageChanger interface {
	ChangeStage(taskId domain.TaskId, stage domain.Stage,
		ensureTaskOwnership domain.TaskOwnershipParams) error
}
