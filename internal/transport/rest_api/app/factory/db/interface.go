package db

import (
	"tracker_backend/internal/application/task"
)

type DbGateway interface {
	task.TaskSaver
	task.OwnerTaskQuerier
	task.TaskStageChanger
}
