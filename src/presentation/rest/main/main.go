package main

import (
	"log"
	"net/http"
	"tracker_backend/src/factory/task"
	"tracker_backend/src/factory/user"
	"tracker_backend/src/infrastructure"
	"tracker_backend/src/presentation/rest"
	"tracker_backend/src/presentation/rest/task_controller"
	"tracker_backend/src/presentation/rest/user_controller"
)

func main() {
	logger := infrastructure.PrintLogger{}

	userSaverFactory := user.UserSaverFactory{}
	createUserFactory := user.CreateUserFactory{
		SaverFactory: userSaverFactory,
	}
	userHandler := user_controller.UserHandler{
		CreateUserFactory: createUserFactory,
		Logger:            logger,
	}

	taskSaverFactory := task.TaskSaverFactory{}
	createTaskFactory := task.CreateFactory{
		SaverFactory: taskSaverFactory,
	}

	stageChangerFactory := task.StageChangerFactory{}
	changeStageFactory := task.ChangeStageFactory{
		StageChangerFactory: stageChangerFactory,
	}

	taskDbGatewayFactory := task.DbQueryGatewayFactory{}
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
	log.Fatal(http.ListenAndServe(":8080", apiHandler))
}
