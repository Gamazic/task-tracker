package db

import (
	"tracker_backend/internal/storage_adapter/task_repo"
	"tracker_backend/internal/transport/rest_api/app/factory"
)

type InMemoryFactory struct {
	inmemoryDb *task_repo.Db
}

func (i *InMemoryFactory) Build(factory.CtxDeps) (DbGateway, error) {
	if i.inmemoryDb == nil {
		i.inmemoryDb = task_repo.NewDb()
	}
	return i.inmemoryDb, nil
}
