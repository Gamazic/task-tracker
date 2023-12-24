package task

import (
	taskUsecase "tracker_backend/src/application/task"
	"tracker_backend/src/presentation/rest_api/app/factory"
	"tracker_backend/src/presentation/rest_api/task_controller"
)

type CreateFactory struct {
	SaverFactory      AbsTaskSaverFactory
	IdProviderFactory task_controller.AbsIdProviderFactory
}

func (c CreateFactory) Build(deps task_controller.CredentialCtxDeps) (task_controller.CreateTaskUsecase, error) {
	saver, err := c.SaverFactory.Build(factory.CtxDeps{Ctx: deps.Ctx})
	if err != nil {
		return nil, err
	}
	idProvider, err := c.IdProviderFactory.Build(deps)
	if err != nil {
		return nil, err
	}
	return taskUsecase.CreateTaskCmd{
		IdProvider: idProvider,
		Saver:      saver,
	}, nil
}
