package task

import (
	taskUsecase "tracker_backend/src/application/task"
	"tracker_backend/src/presentation/rest/app/factory"
	"tracker_backend/src/presentation/rest/task_controller"
)

type GetOwnerTasksFactory struct {
	DbGatewayFactory  AbsDbQueryGatewayFactory
	IdProviderFactory task_controller.AbsIdProviderFactory
}

func (c GetOwnerTasksFactory) Build(deps task_controller.CredentialCtxDeps) (task_controller.GetOwnerTasksUsecase, error) {
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
