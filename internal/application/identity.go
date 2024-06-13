package application

import (
	"errors"
	"fmt"
	"tracker_backend/internal/domain/permission"
)

var (
	ErrProvidingId    = errors.New("failed provide identity")
	ErrNoSuchIdentity = fmt.Errorf("%w: no such identity", ErrProvidingId)
	ErrWrongData      = fmt.Errorf("%w: wrong data", ErrProvidingId)
)

type IdentityProvider interface {
	Provide() (permission.UserRoleIdentity, error)
}
