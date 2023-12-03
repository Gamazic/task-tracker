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

type TaskSaverStubFactory struct{}

func (t TaskSaverStubFactory) Build(deps TaskSaverDeps) (task_command.TaskSaver, error) {
	return taskAdapter.TaskDbGatewayStub{}, nil
}

type StageChangerDeps struct {
	Ctx context.Context
}

type AbsStageChangerFactory interface {
	Build(StageChangerDeps) (task_command.TaskStageChanger, error)
}

type StageChangerStubFactory struct{}

func (t StageChangerStubFactory) Build(deps StageChangerDeps) (task_command.TaskStageChanger, error) {
	return taskAdapter.TaskDbGatewayStub{}, nil
}

type DbQueryGatewayDeps struct {
	Ctx context.Context
}

type AbsDbQueryGatewayFactory interface {
	Build(DbQueryGatewayDeps) (task_query.DbQueryGateway, error)
}

type DbQueryGatewayStubFactory struct{}

func (d DbQueryGatewayStubFactory) Build(DbQueryGatewayDeps) (task_query.DbQueryGateway, error) {
	return taskAdapter.TaskDbGatewayStub{}, nil
}
