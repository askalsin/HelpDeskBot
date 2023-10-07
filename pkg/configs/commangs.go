package configs

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
)

var (
	Commands          CommandsConfig
	onceCommandConfig sync.Once
)

type CommandsConfig struct {
	Start      string `yaml:"start"`
	StartGroup string `yaml:"start_group"`
}

// Возвращает адрес структуры, в которой хранится конфигурация команд бота
func GetComandsConfig() *CommandsConfig {
	onceCommandConfig.Do(func() {
		Commands = CommandsConfig{}
		configFile := "configs/commands.yml"
		configPath := fmt.Sprintf("%s/%s", *GetCurrentWorkingDirectory(), configFile)

		if err := cleanenv.ReadConfig(configPath, &Commands); err != nil {
			log.Error.Fatalf("%s: error parse commands config!", err)
		}
	})

	return &Commands
}
