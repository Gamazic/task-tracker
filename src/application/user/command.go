package user

import (
	"errors"
	"fmt"
	"tracker_backend/src/domain"
)

var ErrCreateUser = errors.New("failed create user")

type UserInCreate struct {
	Username string
}

type CreateUserCmd struct {
	Saver UserSaver
}

func (c CreateUserCmd) Execute(userDto UserInCreate) error {
	user, err := domain.NewUser(userDto.Username)
	if err != nil {
		return err
	}
	err = c.Saver.SaveIfNotExist(user)
	if errors.Is(err, ErrUserAlreadyExist) {
		return fmt.Errorf("create user: %w", err)
	}
	if err != nil {
		return fmt.Errorf("%w: %s", ErrCreateUser, err)
	}
	return nil
}
