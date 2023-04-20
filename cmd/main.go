package main

import (
	"log"

	"github.com/begenov/TaskFlow/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalln(err)
		return
	}
}
