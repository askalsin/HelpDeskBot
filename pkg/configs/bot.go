package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
)

type Token struct {
	Token string `env:"UTELBOTTOKEN"`
}

// Возвращает токен бота, полученный из переменной окружения "UTELBOTTOKEN"
func GetToken() string {
	t := Token{}
	if err := cleanenv.ReadEnv(&t); err != nil {
		log.Error.Fatalf("%s: error parse telegram token!", err)
	}
	return t.Token
}
