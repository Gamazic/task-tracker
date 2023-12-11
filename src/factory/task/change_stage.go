package task

import (
	taskUsecase "tracker_backend/src/application/task"
	"tracker_backend/src/factory"
)

type AbsChangeStageFactory interface {
	Build(deps factory.CredentialCtxDeps) (taskUsecase.ChangeStageUsecase, error)
}

type ChangeStageFactory struct {
	StageChangerFactory AbsStageChangerFactory
	IdProviderFactory   factory.AbsIdProviderFactory
}

func (c ChangeStageFactory) Build(deps factory.CredentialCtxDeps) (taskUsecase.ChangeStageUsecase, error) {
	changer, err := c.StageChangerFactory.Build(factory.CtxDeps{Ctx: deps.Ctx})
	if err != nil {
		return nil, err
	}
	idProvider, err := c.IdProviderFactory.Build(deps)
	if err != nil {
		return nil, err
	}
	return taskUsecase.ChangeTaskStageCmd{
		IdProvider:   idProvider,
		StageChanger: changer,
	}, nil
}
