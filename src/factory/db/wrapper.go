package db

import (
	"tracker_backend/src/application/task"
	"tracker_backend/src/application/user"
	"tracker_backend/src/factory"
)

type UserSaverWrapper struct {
	GatewayFactory AbsDbGatewayFactory
}

func (u UserSaverWrapper) Build(deps factory.CtxDeps) (user.SaveUserUsecase, error) {
	return u.GatewayFactory.Build(deps)
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
