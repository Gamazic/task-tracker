package task_query

import (
	"errors"
	"fmt"
	"tracker_backend/src/application"
	"tracker_backend/src/domain"
)

var (
	ErrGetOwnerTasks = errors.New("failed to get tasks")
)

type OwnerTasksQuery struct {
	OwnerUsername string
	Pagination    application.Pagination
}

type OwnerTasksResult struct {
	OwnerUsername string
	Tasks         []TaskResult
}

type TaskResult struct {
	TaskId      int
	Description string
	Stage       string
}

type GetOwnerTasks struct {
	permService    domain.PermissionService
	requesterRoles domain.UserRoleParams
	db             DbQueryGateway
}

func NewGetOwnerTasks(
	requesterRoles domain.UserRoleParams,
	db DbQueryGateway,
) GetOwnerTasks {
	return GetOwnerTasks{
		permService:    domain.PermissionService{},
		requesterRoles: requesterRoles,
		db:             db,
	}
}

func (g GetOwnerTasks) Execute(query OwnerTasksQuery) (OwnerTasksResult, error) {
	tasksOwnership := domain.TaskOwnershipParams{
		TaskOwnerUsername: domain.Username(query.OwnerUsername),
	}
	if !g.permService.CanRead(g.requesterRoles, tasksOwnership) {
		return OwnerTasksResult{}, domain.ErrOpNotAllowed
	}
	tasks, err := g.db.FetchOwnerTasks(query)
	if err != nil {
		return OwnerTasksResult{}, fmt.Errorf("%w: %s", ErrGetOwnerTasks, err)
	}
	res := OwnerTasksResult{
		OwnerUsername: query.OwnerUsername,
		Tasks:         tasks,
	}
	return res, nil
}
