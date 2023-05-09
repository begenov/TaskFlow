package main

import (
	"fmt"
	"log"

	"github.com/begenov/TaskFlow/pkg/postgresql"
	"github.com/begenov/TaskFlow/task-app/internal/config"
	"github.com/begenov/TaskFlow/task-app/internal/controller"
	"github.com/begenov/TaskFlow/task-app/internal/service"
	"github.com/begenov/TaskFlow/task-app/internal/storage"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalln(err)
		return
	}

	db, err := postgresql.NewDB(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		log.Fatalln(err)
		return
	}

	storage := storage.NewStorage(db)

	service := service.NewService(*storage)

	controller := controller.NewController(service.Task)

	init := controller.Init()

	fmt.Println(init.Run(":8000"))
}
