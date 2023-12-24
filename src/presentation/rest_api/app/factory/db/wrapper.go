package db

import (
	"tracker_backend/src/application/task"
	"tracker_backend/src/presentation/rest_api/app/factory"
)

type AbsDbGatewayFactory interface {
	Build(factory.CtxDeps) (DbGateway, error)
}

type UserSaverWrapper struct {
	GatewayFactory AbsDbGatewayFactory
}

type TaskSaverWrapper struct {
	GatewayFactory AbsDbGatewayFactory
}

func (t TaskSaverWrapper) Build(deps factory.CtxDeps) (task.TaskSaver, error) {
	return t.GatewayFactory.Build(deps)
}

type StageChangerWrapper struct {
	GatewayFactory AbsDbGatewayFactory
}

func (s StageChangerWrapper) Build(deps factory.CtxDeps) (task.TaskStageChanger, error) {
	return s.GatewayFactory.Build(deps)
}

type DbQueryGatewayWrapper struct {
	GatewayFactory AbsDbGatewayFactory
}

func (d DbQueryGatewayWrapper) Build(deps factory.CtxDeps) (task.OwnerTaskQuerier, error) {
	return d.GatewayFactory.Build(deps)
}
