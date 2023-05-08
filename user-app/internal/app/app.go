package app

import (
	"flag"

	"github.com/begenov/TaskFlow/user-app/internal/config"
	"github.com/begenov/TaskFlow/user-app/internal/controller"
	"github.com/begenov/TaskFlow/user-app/internal/service"
	"github.com/begenov/TaskFlow/user-app/internal/storage"
	"github.com/begenov/TaskFlow/user-app/internal/storage/mysql"
)

var dsn *string

const driver = "mysql"

func init() {
	dsn = flag.String("dsn", "task:pass@/task?parseTime=true", "Name mysql")
}

func Run() error {

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	db, err := mysql.NewDB(driver, *dsn)
	if err != nil {
		return err
	}
	storage := storage.NewStorage(db)
	service := service.NewService(*storage, cfg)
	controller := controller.NewController(*service)

	return controller.Router().Run()
}
