package utelbot

import (
	"codeberg.org/kalsin/UtelBot/pkg/types"
	"codeberg.org/kalsin/UtelBot/pkg/yandextacker"
	tgbotapi "github.com/mazanur/telegram-bot-api/v6"
)

var (
	users  map[int64]*types.User
	orders map[int64]*types.Order
	issues map[int64]*yandextacker.RequestIssue
)

func init() {
	users = make(map[int64]*types.User)
	orders = make(map[int64]*types.Order)
	issues = make(map[int64]*yandextacker.RequestIssue)
}

func clearChat(bot *tgbotapi.BotAPI, ChatID int64, lastMessageID int) {
	for i := lastMessageID; i > 0; i-- {
		go bot.Send(tgbotapi.NewDeleteMessage(ChatID, i))
	}
}

//Основная функция для обработки обновлений
func (b *Bot) handle() {
	bot := b.botAPI
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			if !update.FromChat().IsPrivate() {
				continue
			}
			go b.handleCallbackQuery(update.CallbackQuery)
			continue
		}

		if update.Message.IsCommand() {
			if !update.FromChat().IsPrivate() {
				go b.handleCommandGroup(&update)
				continue
			}
			go b.handleCommandPrivate(&update)
			continue
		}

		if update.Message != nil {
			if !update.FromChat().IsPrivate() {
				continue
			}
			go b.handleMessage(update.Message)
			continue
		}
	}
}
