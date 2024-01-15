package db

import (
	"database/sql"
	"tracker_backend/src/adapter/task_repo"
	"tracker_backend/src/presentation/rest_api/app/factory"
)

type PgFactory struct {
	DbName    string
	UserTable string
	TaskTable string
	ConnPool  *sql.DB
}

func (m *PgFactory) Build(deps factory.CtxDeps) (DbGateway, error) {
	return &task_repo.PgDbAdapter{
		TaskTable: m.TaskTable,
		ConnPool:  m.ConnPool,
		Ctx:       deps.Ctx,
	}, nil
}
