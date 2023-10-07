package utelbot

import (
	"strings"

	"codeberg.org/kalsin/UtelBot/pkg/buttons"
	"codeberg.org/kalsin/UtelBot/pkg/configs"
	"codeberg.org/kalsin/UtelBot/pkg/database"
	"codeberg.org/kalsin/UtelBot/pkg/types"
	tgbotapi "github.com/mazanur/telegram-bot-api/v6"
)

// Функция обработки текстовых сообщений
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	bot := b.botAPI
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.ParseMode = "Markdown"

	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()

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

	_, ok = orders[message.Chat.ID]
	if !ok {
		orders[message.Chat.ID] = new(types.Order)
	}

	bot.Send(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))

	if !user.Registered {
		switch user.Stage {
		case stage.Registration.EnterUserName:
			msg = enterOrganization(message)
		case stage.Registration.EnterOrganization:
			msg = enterAddressOrganization(message)
		case stage.Registration.EnterAddressOrganization:
			msg = selectPhoneNumberInputMethod(message)
		case stage.Registration.EnterPhoneNumberTelegramApi:
			msg = enterEmail(message)
		case stage.Registration.EnterPhoneNumberManualy:
			msg = phoneNumberValidating(message)
		case stage.Registration.EnterEmail:
			msg = completeRegistration(message)
		default:
			msg = enterUserName(message.Chat.ID)
		}

		bot.Send(msg)
		return
	}

	splitStage := strings.Split(user.Stage, "*")
	switch splitStage[0] {
	case stage.Menu.MakeOrederText:
		msg = makeOrderCheckData(message)
	case stage.Menu.MakeOrderNewAddress:
		msg = makeOrderCheckData(message)
	case stage.Settings.ChangeName:
		msg = changeName(message)
	case stage.Settings.ChangeOrganization:
		msg = changeOrganization(message)
	case stage.Settings.ChangeAddress:
		msg = changeAddress(message)
	case stage.Settings.ChangeEmail:
		msg = changeEmail(message)
	case stage.Settings.EnterPhoneNumberTelegramApi:
		msg = changePhoneNumber(message)
	case stage.Settings.EnterPhoneNumberManualy:
		msg = phoneNumberValidating(message)
	case stage.Menu.MakeOrderChangeNameContactPerson:
		msg = changeContactPersonPhoneNumber(message)
	case stage.Menu.MakeOrderChangeNumberContactPerson:
		msg = changeContactPersonValidate(message)
	case stage.Menu.SelectionText:
		msg = sendTextAssessment(message)
	default:
		msg.Text = messages.WrongCommand
		msg.ReplyMarkup = buttons.MainMenu()
	}

	bot.Send(msg)
}
