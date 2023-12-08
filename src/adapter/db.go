package adapter

import (
	"tracker_backend/src/application/task"
	"tracker_backend/src/application/user"
)

type DbGateway interface {
	user.SaveUserUsecase
	task.TaskSaver
	task.OwnerTaskQuerier
	task.TaskStageChanger
}
