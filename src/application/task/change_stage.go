package task

import (
	"errors"
	"fmt"
	"tracker_backend/src/application"
	"tracker_backend/src/domain/permission"
	"tracker_backend/src/domain/task"
)

var (
	ErrChangeTaskStage = errors.New("failed to change task's stage")
	ErrTaskDoesntExist = fmt.Errorf("%w: task does not exist", ErrChangeTaskStage)
)

type TaskInStageChange struct {
	TaskNumber  int
	TargetStage string
}

type ChangeTaskStageCmd struct {
	IdProvider   application.IdentityProvider
	StageChanger TaskStageChanger

	permService permission.PermissionService
}

func (c ChangeTaskStageCmd) Execute(taskDto TaskInStageChange) error {
	err := task.TaskNumber(taskDto.TaskNumber).Validate()
	if err != nil {
		return err
	}
	err = task.Stage(taskDto.TargetStage).Validate()
	if err != nil {
		return err
	}
	rolesIdentity, err := c.IdProvider.Provide()
	if errors.Is(err, application.ErrProvidingId) {
		return err
	}
	if err != nil {
		return fmt.Errorf("%w: %s", ErrChangeTaskStage, err)
	}
	taskExist, ownershipMatched, err := c.StageChanger.ChangeStage(ChangeStageDto{
		TaskOwnerUsername: string(rolesIdentity.Username),
		TaskNumber:        taskDto.TaskNumber,
		TargetStage:       taskDto.TargetStage,
	})
	if !taskExist {
		return ErrTaskDoesntExist
	}
	if !ownershipMatched {
		return permission.ErrOpNotAllowed
	}
	if err != nil {
		return fmt.Errorf("%w: %s", ErrChangeTaskStage, err)
	}
	return nil
}
