package task

import (
	"context"
	taskAdapter "tracker_backend/src/adapter/task"
	"tracker_backend/src/application/task/task_command"
	"tracker_backend/src/application/task/task_query"
)

type TaskSaverDeps struct {
	Ctx context.Context
}

type AbsTaskSaverFactory interface {
	Build(TaskSaverDeps) (task_command.TaskSaver, error)
}

type TaskSaverFactory struct{}

func (t TaskSaverFactory) Build(deps TaskSaverDeps) (task_command.TaskSaver, error) {
	return taskAdapter.TaskDbGatewayStub{}, nil
}

type StageChangerDeps struct {
	Ctx context.Context
}

type AbsStageChangerFactory interface {
	Build(StageChangerDeps) (task_command.TaskStageChanger, error)
}

type StageChangerFactory struct{}

func (t StageChangerFactory) Build(deps StageChangerDeps) (task_command.TaskStageChanger, error) {
	return taskAdapter.TaskDbGatewayStub{}, nil
}

type DbQueryGatewayDeps struct {
	Ctx context.Context
}

type AbsDbQueryGatewayFactory interface {
	Build(DbQueryGatewayDeps) (task_query.DbQueryGateway, error)
}

type DbQueryGatewayFactory struct{}

func (d DbQueryGatewayFactory) Build(DbQueryGatewayDeps) (task_query.DbQueryGateway, error) {
	return taskAdapter.TaskDbGatewayStub{}, nil
}
