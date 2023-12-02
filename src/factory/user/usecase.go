package user

import (
	"context"
	"tracker_backend/src/application/user"
)

type CreateUserDeps struct {
	Ctx context.Context
}

type AbsCreateUserFactory interface {
	Build(CreateUserDeps) (user.UserCreator, error)
}

type CreateUserFactory struct {
	SaverFactory AbsUserSaverFactory
}

func (c CreateUserFactory) Build(deps CreateUserDeps) (user.UserCreator, error) {
	saver, err := c.SaverFactory.Build(UserSaverDeps{ctx: deps.Ctx})
	if err != nil {
		return user.CreateUserCmd{}, err
	}
	return user.CreateUserCmd{Saver: saver}, nil
}
