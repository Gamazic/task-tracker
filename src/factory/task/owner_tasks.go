package task

import (
	taskUsecase "tracker_backend/src/application/task"
	"tracker_backend/src/factory"
)

type AbsGetOwnerTasksFactory interface {
	Build(deps factory.CredentialCtxDeps) (taskUsecase.GetOwnerTasksUsecase, error)
}

type GetOwnerTasksFactory struct {
	DbGatewayFactory  AbsDbQueryGatewayFactory
	IdProviderFactory factory.AbsIdProviderFactory
}

func (c GetOwnerTasksFactory) Build(deps factory.CredentialCtxDeps) (taskUsecase.GetOwnerTasksUsecase, error) {
	dbGateway, err := c.DbGatewayFactory.Build(factory.CtxDeps{Ctx: deps.Ctx})
	if err != nil {
		return nil, err
	}
	idProvider, err := c.IdProviderFactory.Build(deps)
	if err != nil {
		return nil, err
	}
	return taskUsecase.GetOwnerTasks{
		IdProvider: idProvider,
		Db:         dbGateway,
	}, nil
}
