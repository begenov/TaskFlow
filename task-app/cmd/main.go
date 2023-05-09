package main

import (
	"log"

	"github.com/begenov/TaskFlow/pkg/postgresql"
	"github.com/begenov/TaskFlow/task-app/internal/config"
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
}
