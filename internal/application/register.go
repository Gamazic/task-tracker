package application

import "errors"

var ErrIdentityAlreadyExist = errors.New("failed to register already exist identity")

type IdentityRegister interface {
	Register() error
}
