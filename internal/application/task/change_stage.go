package task

import (
	"errors"
	"fmt"
	"tracker_backend/internal/application"
	"tracker_backend/internal/domain/permission"
	"tracker_backend/internal/domain/task"
	"tracker_backend/internal/domain/user"
)

var (
	ErrChangeTaskStage = errors.New("failed to change task's stage")
	ErrTaskDoesntExist = fmt.Errorf("%w: task does not exist", ErrChangeTaskStage)
)

type TaskInStageChange struct {
	TaskNumber    int
	TargetStage   string
	OwnerUsername string
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
	ownerUsername := user.Username(taskDto.OwnerUsername)
	err = ownerUsername.Validate()
	if err != nil {
		return err
	}
	taskOwnerships := permission.TaskOwnershipParams{TaskOwnerUsername: ownerUsername}
	requesterRolesIdentity, err := c.IdProvider.Provide()
	if errors.Is(err, application.ErrProvidingId) {
		return err
	}
	if err != nil {
		return fmt.Errorf("%w: %s", ErrChangeTaskStage, err)
	}
	canChange := c.permService.HaveAccess(requesterRolesIdentity, taskOwnerships)
	if !canChange {
		return permission.ErrOpNotAllowed
	}
	taskExist, err := c.StageChanger.ChangeStage(ChangeStageDto{
		TaskOwnerUsername: string(requesterRolesIdentity.Username),
		TaskNumber:        taskDto.TaskNumber,
		TargetStage:       taskDto.TargetStage,
	})
	if !taskExist {
		return ErrTaskDoesntExist
	}
	if err != nil {
		return fmt.Errorf("%w: %s", ErrChangeTaskStage, err)
	}
	return nil
}
