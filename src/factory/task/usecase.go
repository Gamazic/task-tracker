package task

import (
	"context"
	"tracker_backend/src/application/task"
	"tracker_backend/src/application/task/task_command"
	"tracker_backend/src/application/task/task_query"
	"tracker_backend/src/domain"
)

type CreateDeps struct {
	Ctx context.Context
}

type AbsCreateFactory interface {
	Build(deps CreateDeps) (task.TaskCreator, error)
}

type CreateFactory struct {
	SaverFactory AbsTaskSaverFactory
}

func (c CreateFactory) Build(deps CreateDeps) (task.TaskCreator, error) {
	saver, err := c.SaverFactory.Build(TaskSaverDeps{Ctx: deps.Ctx})
	if err != nil {
		return nil, err
	}
	return task_command.CreateTaskCmd{Saver: saver}, nil
}

type ChangeStageDeps struct {
	Ctx               context.Context
	RequesterUsername string
}

type AbsChangeStageFactory interface {
	Build(deps ChangeStageDeps) (task.TaskStageChanger, error)
}

type ChangeStageFactory struct {
	StageChangerFactory AbsStageChangerFactory
}

func (c ChangeStageFactory) Build(deps ChangeStageDeps) (task.TaskStageChanger, error) {
	changer, err := c.StageChangerFactory.Build(StageChangerDeps{Ctx: deps.Ctx})
	if err != nil {
		return nil, err
	}
	requesterRoles, err := domain.NewUserRoleParams(deps.RequesterUsername)
	if err != nil {
		return nil, err
	}
	return task_command.NewChangeTaskStageCmd(requesterRoles, changer), nil
}

type GetOwnerTasksDeps struct {
	Ctx               context.Context
	RequesterUsername string
}

type AbsGetOwnerTasksFactory interface {
	Build(deps GetOwnerTasksDeps) (task.TaskQuerier, error)
}

type GetOwnerTasksFactory struct {
	DbGatewayFactory AbsDbQueryGatewayFactory
}

func (c GetOwnerTasksFactory) Build(deps GetOwnerTasksDeps) (task.TaskQuerier, error) {
	dbGateway, err := c.DbGatewayFactory.Build(DbQueryGatewayDeps{Ctx: deps.Ctx})
	if err != nil {
		return nil, err
	}
	requesterRoles, err := domain.NewUserRoleParams(deps.RequesterUsername)
	if err != nil {
		return nil, err
	}
	return task_query.NewGetOwnerTasks(requesterRoles, dbGateway), nil
}
