package db

import (
	"database/sql"
	"tracker_backend/src/adapter"
	"tracker_backend/src/adapter/pg"
	"tracker_backend/src/factory"
)

type PgFactory struct {
	PgUrl     string
	DbName    string
	UserTable string
	TaskTable string
	ConnPool  *sql.DB
}

func (m *PgFactory) Build(deps factory.CtxDeps) (adapter.DbGateway, error) {
	return &pg.PgDbAdapter{
		TaskTable: m.TaskTable,
		ConnPool:  m.ConnPool,
		Ctx:       deps.Ctx,
	}, nil
}
