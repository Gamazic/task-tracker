package task

import (
	taskUsecase "tracker_backend/src/application/task"
	"tracker_backend/src/presentation/rest_api/app/factory"
	"tracker_backend/src/presentation/rest_api/task_controller"
)

type ChangeStageFactory struct {
	StageChangerFactory AbsStageChangerFactory
	IdProviderFactory   task_controller.AbsIdProviderFactory
}

func (c ChangeStageFactory) Build(deps task_controller.CredentialCtxDeps) (task_controller.ChangeStageUsecase, error) {
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
