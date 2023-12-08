package permission

import (
	"fmt"
	"tracker_backend/src/domain"
	"tracker_backend/src/domain/user"
)

var (
	ErrOpNotAllowed = fmt.Errorf("%w: operation is not allowed", domain.ErrDomain)
)

type UserRoleParams struct {
	Username user.Username
	// ...
	// Group Group
	// Org Org
	// ...
}

type TaskOwnershipParams struct {
	TaskOwnerUsername user.Username
	// ...
	// GroupOwners []Group
	// ...
}
