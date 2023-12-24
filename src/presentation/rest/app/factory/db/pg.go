package db

import (
	"database/sql"
	"tracker_backend/src/adapter/pg"
	"tracker_backend/src/presentation/rest/app/factory"
)

type PgFactory struct {
	DbName    string
	UserTable string
	TaskTable string
	ConnPool  *sql.DB
}

func (m *PgFactory) Build(deps factory.CtxDeps) (DbGateway, error) {
	return &pg.PgDbAdapter{
		TaskTable: m.TaskTable,
		ConnPool:  m.ConnPool,
		Ctx:       deps.Ctx,
	}, nil
}
