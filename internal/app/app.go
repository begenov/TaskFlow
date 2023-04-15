package app

import (
	"github.com/begenov/TaskFlow/internal/pkg/lib/e"
	"github.com/begenov/TaskFlow/internal/storage"
)

func Run() error {
	storage, err := storage.NewStorage()
	if err != nil {
		return e.Wrap("", err)
	}
}
