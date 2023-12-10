package factory

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"tracker_backend/src/adapter/identity"
	"tracker_backend/src/application"
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

type BasicMysqlProviderFactory struct {
	UsersTable string
	MysqlDsn   string
	connPool   *sql.DB
}

func (b *BasicMysqlProviderFactory) Build(deps CredentialCtxDeps) (*identity.BasicAuthMysqlProvider, error) {
	if b.connPool == nil {
		connPool, err := sql.Open("mysql", b.MysqlDsn)
		if err != nil {
			return nil, err
		}
		b.connPool = connPool
	}
	return &identity.BasicAuthMysqlProvider{
		Username:   deps.Username,
		Password:   deps.Password,
		UsersTable: b.UsersTable,
		ConnPool:   b.connPool,
		Ctx:        deps.Ctx,
	}, nil
}

type MysqlIdProviderFactory struct {
	BasicMysqlProviderFactory
}

func (m *MysqlIdProviderFactory) Build(deps CredentialCtxDeps) (application.IdentityProvider, error) {
	return m.BasicMysqlProviderFactory.Build(deps)
}

type MysqlRegisterFactory struct {
	BasicMysqlProviderFactory
}

func (m *MysqlRegisterFactory) Build(deps CredentialCtxDeps) (application.IdentityRegister, error) {
	return m.BasicMysqlProviderFactory.Build(deps)
}
