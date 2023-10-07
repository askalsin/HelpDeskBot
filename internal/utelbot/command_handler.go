package utelbot

import (
	"codeberg.org/kalsin/UtelBot/pkg/buttons"
	"codeberg.org/kalsin/UtelBot/pkg/configs"
	"codeberg.org/kalsin/UtelBot/pkg/database"
	"codeberg.org/kalsin/UtelBot/pkg/types"
	tgbotapi "github.com/mazanur/telegram-bot-api/v6"
)

// Функция обработки введенных команд
func (b *Bot) handleCommandPrivate(update *tgbotapi.Update) {
	bot := b.botAPI
	message := update.Message
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	commands := configs.GetComandsConfig()
	messages := configs.GetMessagesConfig()

	user, ok := users[message.Chat.ID]
	if !ok {
		if usr := database.GetUserData(message.Chat.ID); usr != nil {
			users[message.Chat.ID] = usr
			user = users[message.Chat.ID]
			user.Registered = true
		} else {
			users[message.Chat.ID] = new(types.User)
			user = users[message.Chat.ID]
			user.Registered = false
		}
	}

	if !user.Registered {
		switch message.Command() {
		case commands.Start:
			msg = enterUserName(message.Chat.ID)
		}

		msg.ParseMode = "Markdown"
		bot.Send(msg)
		return
	}

	switch message.Command() {
	case commands.Start:
		msg.Text = messages.SelectAction
		msg.ReplyMarkup = buttons.MainMenu()
	default:
		msg.Text = messages.WrongCommand
		msg.ReplyMarkup = buttons.MainMenu()
	}

	msg.ParseMode = "Markdown"
	bot.Send(msg)
}

func (b *Bot) handleCommandGroup(update *tgbotapi.Update) {
	bot := b.botAPI
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	commands := configs.GetComandsConfig()

	switch update.Message.Command() {
	case commands.StartGroup:
		msg = startGroup(update)
	}

	msg.ParseMode = "Markdown"
	bot.Send(msg)
}
