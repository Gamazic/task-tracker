package user

import (
	"context"
	userAdapter "tracker_backend/src/adapter/user"
	"tracker_backend/src/application/user"
)

type UserSaverDeps struct {
	ctx context.Context
}

type AbsUserSaverFactory interface {
	Build(UserSaverDeps) (user.UserSaver, error)
}

type UserSaverStubFactory struct{}

func (u UserSaverStubFactory) Build(UserSaverDeps) (user.UserSaver, error) {
	return userAdapter.UserSaverStub{}, nil
}
