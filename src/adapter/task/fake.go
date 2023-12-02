package task

import (
	"tracker_backend/src/application/task/task_query"
	"tracker_backend/src/domain"
)

type TaskDbGatewayStub struct{}

func (TaskDbGatewayStub) SaveIncrId(task domain.Task) (domain.TaskId, error) {
	return domain.TaskId(1), nil
}

func (TaskDbGatewayStub) ChangeStage(taskId domain.TaskId, stage domain.Stage, ensureTaskOwnership domain.TaskOwnershipParams) error {
	return nil
}

func (TaskDbGatewayStub) FetchOwnerTasks(query task_query.OwnerTasksQuery) ([]task_query.TaskResult, error) {
	return []task_query.TaskResult{
		{
			TaskId:      1,
			Description: "fake task",
			Stage:       "todo",
		},
	}, nil
}
