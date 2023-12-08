package user

import (
	"errors"
	"fmt"
	userDomain "tracker_backend/src/domain/user"
)

var ErrCreateUser = errors.New("failed create user")
var ErrUserAlreadyExist = errors.New("user already exist")

type UserInCreate struct {
	Username string
}

type CreateUserUsecase struct {
	Saver SaveUserUsecase
}

func (c CreateUserUsecase) Execute(userDto UserInCreate) error {
	user := userDomain.User{
		Username: userDomain.Username(userDto.Username),
	}
	err := user.Validate()
	if err != nil {
		return err
	}
	isFreeUsername, err := c.Saver.SaveCheckFreeUsername(SaverUserDto{Username: string(user.Username)})
	if err != nil {
		return fmt.Errorf("%w: %s", ErrCreateUser, err)
	}
	if !isFreeUsername {
		return ErrUserAlreadyExist
	}
	return nil
}
