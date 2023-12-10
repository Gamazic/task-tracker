package task

import (
	taskUsecase "tracker_backend/src/application/task"
	"tracker_backend/src/factory"
)

type AbsCreateFactory interface {
	Build(deps factory.CredentialCtxDeps) (taskUsecase.CreateTaskUsecase, error)
}

type CreateFactory struct {
	SaverFactory      AbsTaskSaverFactory
	IdProviderFactory factory.AbsIdProviderFactory
}

func (c CreateFactory) Build(deps factory.CredentialCtxDeps) (taskUsecase.CreateTaskUsecase, error) {
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
