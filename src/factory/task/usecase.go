package task

import (
	"context"
	taskUsecase "tracker_backend/src/application/task"
	"tracker_backend/src/domain/permission"
	userDomain "tracker_backend/src/domain/user"
	"tracker_backend/src/factory"
)

type AbsCreateFactory interface {
	Build(deps factory.CtxDeps) (taskUsecase.CreateTaskUsecase, error)
}

type CreateFactory struct {
	SaverFactory AbsTaskSaverFactory
}

func (c CreateFactory) Build(deps factory.CtxDeps) (taskUsecase.CreateTaskUsecase, error) {
	saver, err := c.SaverFactory.Build(deps)
	if err != nil {
		return nil, err
	}
	return taskUsecase.CreateTaskCmd{Saver: saver}, nil
}

type ChangeStageDeps struct {
	Ctx               context.Context
	RequesterUsername string
}

type AbsChangeStageFactory interface {
	Build(deps ChangeStageDeps) (taskUsecase.ChangeStageUsecase, error)
}

type ChangeStageFactory struct {
	StageChangerFactory AbsStageChangerFactory
}

func (c ChangeStageFactory) Build(deps ChangeStageDeps) (taskUsecase.ChangeStageUsecase, error) {
	changer, err := c.StageChangerFactory.Build(factory.CtxDeps{Ctx: deps.Ctx})
	if err != nil {
		return nil, err
	}
	requesterRoles := permission.UserRoleParams{
		Username: userDomain.Username(deps.RequesterUsername),
	}
	if err != nil {
		return nil, err
	}
	return taskUsecase.ChangeTaskStageCmd{
		RequesterRoles: requesterRoles,
		StageChanger:   changer,
	}, nil
}

type GetOwnerTasksDeps struct {
	Ctx               context.Context
	RequesterUsername string
}

type AbsGetOwnerTasksFactory interface {
	Build(deps GetOwnerTasksDeps) (taskUsecase.GetOwnerTasksUsecase, error)
}

type GetOwnerTasksFactory struct {
	DbGatewayFactory AbsDbQueryGatewayFactory
}

func (c GetOwnerTasksFactory) Build(deps GetOwnerTasksDeps) (taskUsecase.GetOwnerTasksUsecase, error) {
	dbGateway, err := c.DbGatewayFactory.Build(factory.CtxDeps{Ctx: deps.Ctx})
	if err != nil {
		return nil, err
	}
	requesterRoles := permission.UserRoleParams{Username: userDomain.Username(deps.RequesterUsername)}
	if err != nil {
		return nil, err
	}
	return taskUsecase.GetOwnerTasks{
		RequesterRoles: requesterRoles,
		Db:             dbGateway,
	}, nil
}
