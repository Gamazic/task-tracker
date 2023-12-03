package main

import (
	"log"
	"net/http"
	"tracker_backend/src/factory"
	"tracker_backend/src/factory/task"
	"tracker_backend/src/factory/user"
	"tracker_backend/src/infrastructure"
	"tracker_backend/src/presentation/rest"
	"tracker_backend/src/presentation/rest/microframework"
	"tracker_backend/src/presentation/rest/task_controller"
	"tracker_backend/src/presentation/rest/user_controller"
)

const bodyMaxBytes = 1024

func main() {
	logger := infrastructure.PrintLogger{}

	inmemoryDbFactory := factory.InMemoryFactory{}

	userSaverFactory := factory.UserSaverWrapper{InMemoryFactory: &inmemoryDbFactory}
	createUserFactory := user.CreateUserFactory{
		SaverFactory: &userSaverFactory,
	}
	userHandler := user_controller.UserHandler{
		CreateUserFactory: createUserFactory,
		Logger:            logger,
	}

	taskSaverFactory := factory.TaskSaverWrapper{InMemoryFactory: &inmemoryDbFactory}
	createTaskFactory := task.CreateFactory{
		SaverFactory: taskSaverFactory,
	}

	stageChangerFactory := factory.StageChangerWrapper{InMemoryFactory: &inmemoryDbFactory}
	changeStageFactory := task.ChangeStageFactory{
		StageChangerFactory: stageChangerFactory,
	}

	taskDbGatewayFactory := factory.DbQueryGatewayWrapper{InMemoryFactory: &inmemoryDbFactory}
	getOwnerTasksFactory := task.GetOwnerTasksFactory{
		DbGatewayFactory: taskDbGatewayFactory,
	}

	taskHandler := task_controller.TaskHandler{
		CreateTaskFactory:    createTaskFactory,
		ChangeStageFactory:   changeStageFactory,
		GetOwnerTasksFactory: getOwnerTasksFactory,
		Logger:               logger,
	}

	apiHandler := rest.MainHandler{
		UserHandler: userHandler,
		TaskHandler: taskHandler,
	}

	mwHandler := microframework.Logging(apiHandler, logger)
	mwHandler = microframework.MaxBytes(mwHandler, bodyMaxBytes)
	log.Fatal(http.ListenAndServe(":8080", mwHandler))
}
