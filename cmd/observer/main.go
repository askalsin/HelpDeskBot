package main

import (
	"codeberg.org/kalsin/UtelBot/internal/observer"
	"codeberg.org/kalsin/UtelBot/pkg/database"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Error.Fatalln(err)
	}
	defer database.Close()

	observer.Start()
}
