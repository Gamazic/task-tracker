package task

import (
	"tracker_backend/src/application/task/task_command"
	"tracker_backend/src/application/task/task_query"
)

type TaskCreator interface {
	Execute(task_command.TaskInCreate) (task_command.CreatedTaskArtefacts, error)
}

type TaskStageChanger interface {
	Execute(change task_command.TaskInStageChange) error
}

type TaskQuerier interface {
	Execute(query task_query.OwnerTasksQuery) (task_query.OwnerTasksResult, error)
}
