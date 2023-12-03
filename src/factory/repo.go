package factory

import (
	"tracker_backend/src/adapter/inmemory"
	"tracker_backend/src/application/task/task_command"
	"tracker_backend/src/application/task/task_query"
	"tracker_backend/src/application/user"
	taskFactory "tracker_backend/src/factory/task"
	userFactory "tracker_backend/src/factory/user"
)

type InMemoryFactory struct {
	inmemoryDb *inmemory.Db
}

func (i *InMemoryFactory) Build() *inmemory.Db {
	if i.inmemoryDb == nil {
		i.inmemoryDb = inmemory.NewDb()
	}
	return i.inmemoryDb
}

type UserSaverWrapper struct {
	InMemoryFactory *InMemoryFactory
}

func (u UserSaverWrapper) Build(deps userFactory.UserSaverDeps) (user.UserSaver, error) {
	return u.InMemoryFactory.Build(), nil
}

type TaskSaverWrapper struct {
	InMemoryFactory *InMemoryFactory
}

func (t TaskSaverWrapper) Build(deps taskFactory.TaskSaverDeps) (task_command.TaskSaver, error) {
	return t.InMemoryFactory.Build(), nil
}

type StageChangerWrapper struct {
	InMemoryFactory *InMemoryFactory
}

func (s StageChangerWrapper) Build(deps taskFactory.StageChangerDeps) (task_command.TaskStageChanger, error) {
	return s.InMemoryFactory.Build(), nil
}

type DbQueryGatewayWrapper struct {
	InMemoryFactory *InMemoryFactory
}

func (d DbQueryGatewayWrapper) Build(deps taskFactory.DbQueryGatewayDeps) (task_query.DbQueryGateway, error) {
	return d.InMemoryFactory.Build(), nil
}
