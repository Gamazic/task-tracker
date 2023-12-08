package db

import (
	"tracker_backend/src/adapter"
	"tracker_backend/src/adapter/inmemory"
	"tracker_backend/src/factory"
)

type InMemoryFactory struct {
	inmemoryDb *inmemory.Db
}

func (i *InMemoryFactory) Build(factory.CtxDeps) (adapter.DbGateway, error) {
	if i.inmemoryDb == nil {
		i.inmemoryDb = inmemory.NewDb()
	}
	return i.inmemoryDb, nil
}
