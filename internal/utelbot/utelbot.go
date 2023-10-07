package utelbot

import (
	tgbotapi "github.com/mazanur/telegram-bot-api/v6"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
)

type Bot struct {
	botAPI *tgbotapi.BotAPI
}

func NewBot(token string) *Bot {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Error.Fatalln(err)
	}

	return &Bot {
		botAPI: botAPI,
	}
}

func (b *Bot) Start() {
	b.botAPI.Debug = false
	log.Info.Printf("authorized on account %s", b.botAPI.Self.UserName)
	b.handle()
}
