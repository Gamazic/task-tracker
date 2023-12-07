package task_command

import (
	"errors"
	"fmt"
	"tracker_backend/src/domain"
)

var (
	ErrChangeTaskStage = errors.New("failed to change task's stage")
)

type TaskInStageChange struct {
	TaskId      int
	TargetStage string
}

type ChangeTaskStageCmd struct {
	requesterRoles domain.UserRoleParams
	permService    domain.PermissionService

	stageChanger TaskStageChanger
}

func NewChangeTaskStageCmd(
	requesterRoles domain.UserRoleParams,
	stageChanger TaskStageChanger,
) ChangeTaskStageCmd {
	return ChangeTaskStageCmd{
		requesterRoles: requesterRoles,
		permService:    domain.PermissionService{},
		stageChanger:   stageChanger,
	}
}

func (c ChangeTaskStageCmd) Execute(taskDto TaskInStageChange) error {
	taskId := domain.TaskId(taskDto.TaskId)
	targetStage, err := domain.NewStage(taskDto.TargetStage)
	if err != nil {
		return err
	}
	requireOwnership := c.permService.UserRoleToTask(c.requesterRoles)
	err = c.stageChanger.ChangeStage(taskId, targetStage, requireOwnership)
	if errors.Is(err, domain.ErrOpNotAllowed) {
		return err
	}
	if errors.Is(err, ErrTaskNotFound) {
		return err
	}
	if err != nil {
		return fmt.Errorf("%w: %s", ErrChangeTaskStage, err)
	}
	return nil
}
