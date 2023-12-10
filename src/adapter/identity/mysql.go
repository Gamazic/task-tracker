package identity

import (
	"bytes"
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"tracker_backend/src/application"
	"tracker_backend/src/domain/permission"
	"tracker_backend/src/domain/user"
)

type BasicAuthMysqlProvider struct {
	Username string
	Password string

	UsersTable string
	ConnPool   *sql.DB

	Ctx context.Context
}

func (b *BasicAuthMysqlProvider) Provide() (permission.UserRoleIdentity, error) {
	query := fmt.Sprintf("SELECT password FROM %s WHERE username = ?;", b.UsersTable)
	rows, err := b.ConnPool.QueryContext(b.Ctx, query, b.Username)
	if err != nil {
		return permission.UserRoleIdentity{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return permission.UserRoleIdentity{}, application.ErrNoSuchIdentity
	}
	var persistPswrdHash []byte
	rows.Scan(&persistPswrdHash)

	if !bytes.Equal(b.pswrdHash(), persistPswrdHash) {
		return permission.UserRoleIdentity{}, application.ErrWrongData
	}
	return permission.UserRoleIdentity{
		Username: user.Username(b.Username),
	}, nil
}

func (b *BasicAuthMysqlProvider) Register() error {
	query := fmt.Sprintf("INSERT IGNORE INTO %s (username, password) VALUES(?, ?);", b.UsersTable)
	res, err := b.ConnPool.ExecContext(b.Ctx, query, b.Username, b.pswrdHash())
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return application.ErrIdentityAlreadyExist
	}
	return nil
}

func (b *BasicAuthMysqlProvider) pswrdHash() []byte {
	hash := md5.Sum([]byte(b.Password))
	return hash[:]
}
