package main

import (
	"log"
	"tracker_backend/src/presentation/rest"
)

func main() {
	app := rest.App{
		PgConf: rest.PgConf{
			Url:       "postgres://root:example@localhost:5432/tasktracker",
			DbName:    "tasktracker",
			UserTable: "user",
			TaskTable: "task",
		},
		SwaggerDirPath:  "./swagger",
		ApiBodyMaxBytes: 1024,
		ApiAddr:         "0.0.0.0:8080",
	}
	log.Fatal(app.Run())
}
