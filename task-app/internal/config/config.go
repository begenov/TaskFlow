package config

import (
	"os"

	"github.com/subosito/gotenv"
)

const path = "../.env"

type ConfigTask struct {
	Database db
}

type db struct {
	Driver string
	DSN    string
}

func NewConfig() (*ConfigTask, error) {
	err := gotenv.Load(path)

	if err != nil {
		return nil, err
	}

	dsn := os.Getenv("DSN")
	driver := os.Getenv("DRIVER")

	return &ConfigTask{
		Database: db{Driver: driver, DSN: dsn},
	}, nil
}
