package configs

import (
	"os"
	"sync"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
)

var (
	onceCurrentWorkingDirectory sync.Once
	currentWorkingDirectory     string
)

// Функция, которая возвращает root дирректорию приложения
func GetCurrentWorkingDirectory() *string {
	onceCurrentWorkingDirectory.Do(func() {
		var err error
		currentWorkingDirectory, err = os.Getwd()
		if err != nil {
			log.Error.Fatal(err)
		}
	})

	return &currentWorkingDirectory
}
