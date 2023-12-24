package app

import (
	"database/sql"
	"log"
	"net/http"
	"tracker_backend/src/infrastructure"
	"tracker_backend/src/presentation/rest"
	"tracker_backend/src/presentation/rest/app/factory"
	"tracker_backend/src/presentation/rest/app/factory/db"
	"tracker_backend/src/presentation/rest/app/factory/task"
	"tracker_backend/src/presentation/rest/microframework"
	"tracker_backend/src/presentation/rest/register_controller"
	"tracker_backend/src/presentation/rest/task_controller"
)

type PgConf struct {
	Url       string
	DbName    string
	UserTable string
	TaskTable string
}

type App struct {
	PgConf          PgConf
	SwaggerDirPath  string
	ApiBodyMaxBytes int64
	ApiAddr         string
}

func (a App) Run() error {
	logger := infrastructure.PrintLogger{}

	connPool, err := sql.Open("pgx", a.PgConf.Url)
	if err != nil {
		log.Fatal(err)
	}
	dbFactory := db.PgFactory{
		DbName:    a.PgConf.DbName,
		UserTable: a.PgConf.UserTable,
		TaskTable: a.PgConf.TaskTable,
		ConnPool:  connPool,
	}
	//dbFactory := db.InMemoryFactory{}

	pgIdFactory := factory.BasicPgProviderFactory{
		UserTable: a.PgConf.UserTable,
		ConnPool:  connPool,
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

	swaggerHandler := http.FileServer(http.Dir(a.SwaggerDirPath))
	apiHandler := rest.MainHandler{
		RegisterController: registerController,
		TaskController:     taskController,
		SwaggerHandler:     swaggerHandler,
	}

	mwHandler := microframework.BasicAuthentication(apiHandler, "/api/tasks")
	mwHandler = microframework.MaxBytes(mwHandler, a.ApiBodyMaxBytes)
	mwHandler = microframework.Logging(mwHandler, logger)
	return http.ListenAndServe(a.ApiAddr, mwHandler)
}
