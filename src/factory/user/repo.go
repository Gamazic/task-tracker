package user

import (
	"tracker_backend/src/application/user"
	"tracker_backend/src/factory"
)

type AbsUserSaverFactory interface {
	Build(deps factory.CtxDeps) (user.SaveUserUsecase, error)
}

//type UserSaverStubFactory struct{}
//
//func (u UserSaverStubFactory) Build(factory.CtxDeps) (user.UserSaver, error) {
//	return userAdapter.UserSaverStub{}, nil
//}
