package adapter

import (
	"tracker_backend/src/application/task"
)

type DbGateway interface {
	task.TaskSaver
	task.OwnerTaskQuerier
	task.TaskStageChanger
}
