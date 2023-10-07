package configs

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
)

var (
	onceDBconfig sync.Once
	DBcfg        DBConfig
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `env:"UTELBOTDBPASSWORD"`
	Port     int16  `yaml:"port"`
}

// Возвращает адрес структуры, в которой хранится конфигурация БД
func GetDBConfig() *DBConfig {
	onceDBconfig.Do(func() {
		DBcfg = DBConfig{}
		configFile := "configs/database.yml"
		configPath := fmt.Sprintf("%s/%s", *GetCurrentWorkingDirectory(), configFile)

		if err := cleanenv.ReadConfig(configPath, &DBcfg); err != nil {
			log.Error.Fatalf("%s: error parse DB configs!", err)
		}

		if err := cleanenv.ReadEnv(&DBcfg); err != nil {
			log.Error.Fatalf("%s: error parse DB password!", err)
		}
	})

	return &DBcfg
}

func (db *DBConfig) ConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		db.User, db.Password, db.Host, db.Port, db.Name)
}
