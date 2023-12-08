package user

import (
	"tracker_backend/src/application/user"
	"tracker_backend/src/factory"
)

type AbsCreateUserFactory interface {
	Build(factory.CtxDeps) (user.UserCreator, error)
}

type CreateUserFactory struct {
	SaverFactory AbsUserSaverFactory
}

func (c CreateUserFactory) Build(deps factory.CtxDeps) (user.UserCreator, error) {
	saver, err := c.SaverFactory.Build(deps)
	if err != nil {
		return user.CreateUserUsecase{}, err
	}
	return user.CreateUserUsecase{Saver: saver}, nil
}
