package register_controller

import (
	"context"
	"tracker_backend/src/application"
)

type RegisterDeps struct {
	Ctx      context.Context
	Username string
	Password string
}

type AbsRegisterFactory interface {
	Build(deps RegisterDeps) (application.IdentityRegister, error)
}
