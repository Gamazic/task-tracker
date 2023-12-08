package task

import (
	"tracker_backend/src/application/task"
	"tracker_backend/src/factory"
)

type AbsTaskSaverFactory interface {
	Build(factory.CtxDeps) (task.TaskSaver, error)
}

//type TaskSaverStubFactory struct{}
//
//func (t TaskSaverStubFactory) Build(deps factory.CtxDeps) (task_command.TaskSaver, error) {
//	return taskAdapter.TaskDbGatewayStub{}, nil
//}

type AbsStageChangerFactory interface {
	Build(factory.CtxDeps) (task.TaskStageChanger, error)
}

//type StageChangerStubFactory struct{}
//
//func (t StageChangerStubFactory) Build(deps factory.CtxDeps) (task_command.TaskStageChanger, error) {
//	return taskAdapter.TaskDbGatewayStub{}, nil
//}

type AbsDbQueryGatewayFactory interface {
	Build(factory.CtxDeps) (task.OwnerTaskQuerier, error)
}

//type DbQueryGatewayStubFactory struct{}
//
//func (d DbQueryGatewayStubFactory) Build(factory.CtxDeps) (task_query.DbQueryGateway, error) {
//	return taskAdapter.TaskDbGatewayStub{}, nil
//}
