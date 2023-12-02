package task_query

type DbQueryGateway interface {
	FetchOwnerTasks(OwnerTasksQuery) ([]TaskResult, error)
}
