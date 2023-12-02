package domain

import "errors"

var (
	ErrOpNotAllowed = errors.New("operation is not allowed")
)

type UserRoleParams struct {
	Username Username
	// ...
	// Group Group
	// Org Org
	// ...
}

func NewUserRoleParams(username string) (UserRoleParams, error) {
	u, err := NewUsername(username)
	if err != nil {
		return UserRoleParams{}, err
	}
	return UserRoleParams{
		Username: u,
	}, nil
}

type TaskOwnershipParams struct {
	TaskOwnerUsername Username
	// ...
	// GroupOwners []Group
	// ...
}

type PermissionService struct{}

func (p PermissionService) CanRead(user UserRoleParams, task TaskOwnershipParams) bool {
	return task.TaskOwnerUsername == user.Username
}

func (p PermissionService) UserRoleToTask(user UserRoleParams) TaskOwnershipParams {
	return TaskOwnershipParams{
		TaskOwnerUsername: user.Username,
	}
}
