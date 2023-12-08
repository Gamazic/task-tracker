package user

import (
	"fmt"
	"tracker_backend/src/domain"
)

var (
	ErrInvalidUsername = fmt.Errorf("%w: invalid username", domain.ErrDomain)
	ErrEmptyUsername   = fmt.Errorf("%w: empty username", ErrInvalidUsername)
)

type Username string

func (u Username) Validate() error {
	if u == "" {
		return ErrEmptyUsername
	}
	return nil
}
