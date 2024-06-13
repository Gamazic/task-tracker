package task

type CreateTaskUsecase interface {
	Execute(TaskInCreate) (CreatedTaskArtefacts, error)
}

type ChangeStageUsecase interface {
	Execute(change TaskInStageChange) error
}

type GetOwnerTasksUsecase interface {
	Execute(query OwnerTasksQuery) ([]TaskResult, error)
}
