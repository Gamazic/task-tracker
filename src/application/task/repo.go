package task

type TaskSaveDto struct {
	Description   string
	OwnerUsername string
	Stage         string
}

type TaskSaver interface {
	SaveIncrOwnerTaskNumber(task TaskSaveDto) (int, error)
}

type OwnerTaskQuerier interface {
	FetchOwnerTasks(OwnerTasksQuery) ([]TaskResult, error)
}

type ChangeStageDto struct {
	TaskOwnerUsername string
	TaskNumber        int
	TargetStage       string
}

type TaskStageChanger interface {
	ChangeStage(ChangeStageDto) (taskExist bool, err error)
}
