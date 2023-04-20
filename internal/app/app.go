package app

import (
	"flag"
	"net/http"

	"github.com/begenov/TaskFlow/internal/controller"
	"github.com/begenov/TaskFlow/internal/service"
	"github.com/begenov/TaskFlow/internal/storage"
	"github.com/begenov/TaskFlow/internal/storage/mysql"
)

var dsn *string

const driver = "mysql"

func init() {
	dsn = flag.String("dsn", "task:pass@/task?parseTime=true", "Name mysql")
}

func Run() error {
	db, err := mysql.NewDB(driver, *dsn)
	if err != nil {
		return err
	}
	storage := storage.NewStorage(db)
	service := service.NewService(*storage)
	controller := controller.NewController(*service)

	return http.ListenAndServe(":8080", controller.Router())
}
