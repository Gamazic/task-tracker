package user

import (
	"errors"
	"tracker_backend/src/domain"
)

var (
	ErrUserAlreadyExist = errors.New("user already exist")
)

type UserSaver interface {
	SaveIfNotExist(user domain.User) error
}
