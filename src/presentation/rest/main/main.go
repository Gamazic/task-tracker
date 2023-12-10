package main

import (
	"log"
	"net/http"
	"tracker_backend/src/factory"
	"tracker_backend/src/factory/db"
	"tracker_backend/src/factory/task"
	"tracker_backend/src/infrastructure"
	"tracker_backend/src/presentation/rest"
	"tracker_backend/src/presentation/rest/microframework"
	"tracker_backend/src/presentation/rest/register_controller"
	"tracker_backend/src/presentation/rest/task_controller"
)

const bodyMaxBytes = 1024

func main() {
	logger := infrastructure.PrintLogger{}

	dbFactory := db.InMemoryFactory{}
	//dbFactory := db.PgFactory{
	//	PgDsn:  "root:example@/tasktracker",
	//	DbName:    "tasktracker",
	//	UserTable: "user",
	//	TaskTable: "task",
	//}

	pgIdFactory := factory.BasicPgProviderFactory{
		UsersTable: "user",
		PgDsn:      "postgres://root:example@localhost:5432/tasktracker",
	}
	pgIdProviderFactory := factory.PgIdProviderFactory{pgIdFactory}
	pgRegisterFactory := factory.PgRegisterFactory{pgIdFactory}

	taskSaverFactory := db.TaskSaverWrapper{GatewayFactory: &dbFactory}
	createTaskFactory := task.CreateFactory{
		SaverFactory:      &taskSaverFactory,
		IdProviderFactory: &pgIdProviderFactory,
	}

	stageChangerFactory := db.StageChangerWrapper{GatewayFactory: &dbFactory}
	changeStageFactory := task.ChangeStageFactory{
		StageChangerFactory: &stageChangerFactory,
		IdProviderFactory:   &pgIdProviderFactory,
	}

	taskDbGatewayFactory := db.DbQueryGatewayWrapper{GatewayFactory: &dbFactory}
	getOwnerTasksFactory := task.GetOwnerTasksFactory{
		DbGatewayFactory:  &taskDbGatewayFactory,
		IdProviderFactory: &pgIdProviderFactory,
	}

	taskController := task_controller.TaskController{
		CreateTaskFactory:    createTaskFactory,
		ChangeStageFactory:   changeStageFactory,
		GetOwnerTasksFactory: getOwnerTasksFactory,
		Logger:               logger,
	}

	registerController := register_controller.RegisterController{
		RegisterFactory: &pgRegisterFactory,
		Logger:          logger,
	}

	swaggerHandler := http.FileServer(http.Dir("./swagger"))
	apiHandler := rest.MainHandler{
		RegisterController: registerController,
		TaskController:     taskController,
		SwaggerHandler:     swaggerHandler,
	}

	mwHandler := microframework.Logging(apiHandler, logger)
	mwHandler = microframework.MaxBytes(mwHandler, bodyMaxBytes)
	mwHandler = microframework.BasicAuthentication(mwHandler, "/api/tasks")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mwHandler))
}
