package task

import (
	"tracker_backend/src/application/task"
	"tracker_backend/src/presentation/rest/app/factory"
)

type AbsTaskSaverFactory interface {
	Build(factory.CtxDeps) (task.TaskSaver, error)
}

type AbsStageChangerFactory interface {
	Build(factory.CtxDeps) (task.TaskStageChanger, error)
}

type AbsDbQueryGatewayFactory interface {
	Build(factory.CtxDeps) (task.OwnerTaskQuerier, error)
}
