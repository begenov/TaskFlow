package app

import (
	"flag"

	"github.com/begenov/TaskFlow/pkg/auth"
	"github.com/begenov/TaskFlow/pkg/postgresql"
	"github.com/begenov/TaskFlow/user-app/internal/config"
	"github.com/begenov/TaskFlow/user-app/internal/controller"
	"github.com/begenov/TaskFlow/user-app/internal/service"
	"github.com/begenov/TaskFlow/user-app/internal/storage"
)

var dsn *string

const driver = "postgres"

func init() {
	dsn = flag.String("dsn", "postgresql://root:secret@localhost:5432/user?sslmode=disable", "Name postgresql")
}

func Run() error {

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	tokenManager, err := auth.NewManager(cfg.JWT.SigninKey)
	if err != nil {
		return err
	}
	db, err := postgresql.NewDB(driver, *dsn)
	if err != nil {
		return err
	}
	storage := storage.NewStorage(db)
	service := service.NewService(*storage, cfg, tokenManager)
	controller := controller.NewController(*service)

	return controller.Router().Run()
}
