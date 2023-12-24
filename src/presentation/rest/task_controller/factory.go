package task_controller

import (
	"context"
	"tracker_backend/src/application"
	taskUsecase "tracker_backend/src/application/task"
)

type CredentialCtxDeps struct {
	Ctx      context.Context
	Username string
	Password string
}

type AbsChangeStageFactory interface {
	Build(deps CredentialCtxDeps) (ChangeStageUsecase, error)
}
type ChangeStageUsecase interface {
	Execute(change taskUsecase.TaskInStageChange) error
}

type AbsCreateFactory interface {
	Build(deps CredentialCtxDeps) (CreateTaskUsecase, error)
}
type CreateTaskUsecase interface {
	Execute(create taskUsecase.TaskInCreate) (taskUsecase.CreatedTaskArtefacts, error)
}

type AbsGetOwnerTasksFactory interface {
	Build(deps CredentialCtxDeps) (GetOwnerTasksUsecase, error)
}
type GetOwnerTasksUsecase interface {
	Execute(query taskUsecase.OwnerTasksQuery) ([]taskUsecase.TaskResult, error)
}

type AbsIdProviderFactory interface {
	Build(deps CredentialCtxDeps) (application.IdentityProvider, error)
}
