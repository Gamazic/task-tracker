package task

import (
	"errors"
	"fmt"
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
	RequesterRoles permission.UserRoleParams
	StageChanger   TaskStageChanger

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
	taskExist, ownershipMatched, err := c.StageChanger.ChangeStage(ChangeStageDto{
		TaskOwnerUsername: string(c.RequesterRoles.Username),
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
