package task

import (
	"tracker_backend/internal/application/task"
	"tracker_backend/internal/transport/rest_api/app/factory"
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
