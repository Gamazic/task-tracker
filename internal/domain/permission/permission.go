package permission

import (
	"fmt"
	"tracker_backend/internal/domain"
	"tracker_backend/internal/domain/user"
)

var (
	ErrOpNotAllowed = fmt.Errorf("%w: operation is not allowed", domain.ErrDomain)
)

type UserRoleIdentity struct {
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
