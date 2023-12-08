package main

import (
	"log"
	"net/http"
	"tracker_backend/src/factory/db"
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

	dbFactory := db.InMemoryFactory{}
	//dbFactory := db.MysqlFactory{
	//	MysqlDsn:  "root:example@/tasktracker",
	//	DbName:    "tasktracker",
	//	UserTable: "user",
	//	TaskTable: "task",
	//}

	userSaverFactory := db.UserSaverWrapper{GatewayFactory: &dbFactory}
	createUserFactory := user.CreateUserFactory{
		SaverFactory: &userSaverFactory,
	}
	userHandler := user_controller.UserHandler{
		CreateUserFactory: createUserFactory,
		Logger:            logger,
	}

	taskSaverFactory := db.TaskSaverWrapper{GatewayFactory: &dbFactory}
	createTaskFactory := task.CreateFactory{
		SaverFactory: &taskSaverFactory,
	}

	stageChangerFactory := db.StageChangerWrapper{GatewayFactory: &dbFactory}
	changeStageFactory := task.ChangeStageFactory{
		StageChangerFactory: &stageChangerFactory,
	}

	taskDbGatewayFactory := db.DbQueryGatewayWrapper{GatewayFactory: &dbFactory}
	getOwnerTasksFactory := task.GetOwnerTasksFactory{
		DbGatewayFactory: &taskDbGatewayFactory,
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
		SwaggerDir:  "./swagger",
	}

	mwHandler := microframework.Logging(apiHandler, logger)
	mwHandler = microframework.MaxBytes(mwHandler, bodyMaxBytes)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mwHandler))
}
