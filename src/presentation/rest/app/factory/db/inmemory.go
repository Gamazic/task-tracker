package db

import (
	"tracker_backend/src/adapter/inmemory"
	"tracker_backend/src/presentation/rest/app/factory"
)

type InMemoryFactory struct {
	inmemoryDb *inmemory.Db
}

func (i *InMemoryFactory) Build(factory.CtxDeps) (DbGateway, error) {
	if i.inmemoryDb == nil {
		i.inmemoryDb = inmemory.NewDb()
	}
	return i.inmemoryDb, nil
}
