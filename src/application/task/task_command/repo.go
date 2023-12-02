package task_command

import "tracker_backend/src/domain"

type TaskSaver interface {
	SaveIncrId(task domain.Task) (domain.TaskId, error)
}

type TaskStageChanger interface {
	ChangeStage(taskId domain.TaskId, stage domain.Stage,
		ensureTaskOwnership domain.TaskOwnershipParams) error
}
