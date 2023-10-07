package utelbot

import (
	"codeberg.org/kalsin/UtelBot/pkg/configs"
	"codeberg.org/kalsin/UtelBot/pkg/database"
	tgbotapi "github.com/mazanur/telegram-bot-api/v6"
)

func contains(sl []int64, val int64) bool {
	for _, v := range sl {
		if v == val {
			return true
		}
	}
	return false
}

func startGroup(update *tgbotapi.Update) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ReplyToMessageID = update.Message.MessageID
	
	chatsID := database.GetGroupsChatID()
	if !contains(chatsID, update.FromChat().ID) {
		database.NewGroup(update.FromChat().ID)
	}
	
	msg.ParseMode = "Markdown"
	msg.Text = messages.StartGroup

	return msg
}
