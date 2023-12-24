package factory

import (
	"database/sql"
	"tracker_backend/src/adapter/identity"
	"tracker_backend/src/application"
	"tracker_backend/src/presentation/rest_api/register_controller"
	"tracker_backend/src/presentation/rest_api/task_controller"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type BasicPgProviderFactory struct {
	UserTable string
	ConnPool  *sql.DB
}

func (b *BasicPgProviderFactory) Build(deps task_controller.CredentialCtxDeps) (*identity.BasicAuthPgProvider, error) {
	return &identity.BasicAuthPgProvider{
		Username:   deps.Username,
		Password:   deps.Password,
		UsersTable: b.UserTable,
		ConnPool:   b.ConnPool,
		Ctx:        deps.Ctx,
	}, nil
}

type PgIdProviderFactory struct {
	BasicPgProviderFactory
}

func (m *PgIdProviderFactory) Build(deps task_controller.CredentialCtxDeps) (application.IdentityProvider, error) {
	return m.BasicPgProviderFactory.Build(deps)
}

type PgRegisterFactory struct {
	BasicPgProviderFactory
}

func (m *PgRegisterFactory) Build(deps register_controller.RegisterDeps) (application.IdentityRegister, error) {
	return m.BasicPgProviderFactory.Build(task_controller.CredentialCtxDeps{
		Ctx:      deps.Ctx,
		Username: deps.Username,
		Password: deps.Password,
	})
}
