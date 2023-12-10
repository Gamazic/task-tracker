package task

import (
	"errors"
	"fmt"
	"tracker_backend/src/application"
	"tracker_backend/src/domain/permission"
	"tracker_backend/src/domain/user"
)

var (
	ErrGetOwnerTasks = errors.New("failed to get tasks")
)

type OwnerTasksQuery struct {
	OwnerUsername string
}

type TaskResult struct {
	TaskNumber  int
	Description string
	Stage       string
}

type GetOwnerTasks struct {
	IdProvider  application.IdentityProvider
	Db          OwnerTaskQuerier
	permService permission.PermissionService
}

func (g GetOwnerTasks) Execute(query OwnerTasksQuery) ([]TaskResult, error) {
	ownerUsername := user.Username(query.OwnerUsername)
	err := ownerUsername.Validate()
	if err != nil {
		return nil, err
	}
	requesterRoles, err := g.IdProvider.Provide()
	if err != nil {
		return nil, err
	}
	tasksOwnership := permission.TaskOwnershipParams{
		TaskOwnerUsername: ownerUsername,
	}
	if !g.permService.HaveAccess(requesterRoles, tasksOwnership) {
		return nil, permission.ErrOpNotAllowed
	}
	tasks, err := g.Db.FetchOwnerTasks(query)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrGetOwnerTasks, err)
	}
	return tasks, nil
}
