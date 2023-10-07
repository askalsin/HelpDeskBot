package configs

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
)

var (
	TrackerConfig     YandexTrackerConfig
	onceTrackerConfig sync.Once
)

type YandexTrackerConfig struct {
	Token string `env:"UTELYANDEXTRACKERTOKEN"`
	OrgID string `env:"YANDEXTRACKERIDORGANIZATION"`
	Queue string `yaml:"tracker_queue"`
}

// Возвращает токен бота, полученный из переменной окружения "UTELBOTTOKEN"
func GetTrackerConfig() *YandexTrackerConfig {
	onceTrackerConfig.Do(func() {
		TrackerConfig = YandexTrackerConfig{}
		if err := cleanenv.ReadEnv(&TrackerConfig); err != nil {
			log.Error.Fatalf("%s: error parse yandex tracker configs with env!", err)
		}
		configFile := "configs/yandextracker.yml"
		configPath := fmt.Sprintf("%s/%s", *GetCurrentWorkingDirectory(), configFile)

		if err := cleanenv.ReadConfig(configPath, &TrackerConfig); err != nil {
			log.Error.Fatalf("%s: error parse yandex tracker configs with yaml file!", err)
		}
	})
	return &TrackerConfig
}
