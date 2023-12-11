package factory

import (
	"context"
	"database/sql"
	"tracker_backend/src/adapter/identity"
	"tracker_backend/src/application"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type CredentialCtxDeps struct {
	Ctx      context.Context
	Username string
	Password string
}

type AbsIdProviderFactory interface {
	Build(deps CredentialCtxDeps) (application.IdentityProvider, error)
}

type AbsRegisterFactory interface {
	Build(deps CredentialCtxDeps) (application.IdentityRegister, error)
}

type BasicPgProviderFactory struct {
	UsersTable string
	ConnPool   *sql.DB
}

func (b *BasicPgProviderFactory) Build(deps CredentialCtxDeps) (*identity.BasicAuthPgProvider, error) {
	return &identity.BasicAuthPgProvider{
		Username:   deps.Username,
		Password:   deps.Password,
		UsersTable: b.UsersTable,
		ConnPool:   b.ConnPool,
		Ctx:        deps.Ctx,
	}, nil
}

type PgIdProviderFactory struct {
	BasicPgProviderFactory
}

func (m *PgIdProviderFactory) Build(deps CredentialCtxDeps) (application.IdentityProvider, error) {
	return m.BasicPgProviderFactory.Build(deps)
}

type PgRegisterFactory struct {
	BasicPgProviderFactory
}

func (m *PgRegisterFactory) Build(deps CredentialCtxDeps) (application.IdentityRegister, error) {
	return m.BasicPgProviderFactory.Build(deps)
}
