package utelbot

import (
	"strconv"
	"strings"

	"codeberg.org/kalsin/UtelBot/pkg/configs"
	"codeberg.org/kalsin/UtelBot/pkg/database"
	"codeberg.org/kalsin/UtelBot/pkg/types"
	tgbotapi "github.com/mazanur/telegram-bot-api/v6"
)

// Функция для обработки нажатий на кнопки
func (b *Bot) handleCallbackQuery(callbackQuery *tgbotapi.CallbackQuery) {
	bot := b.botAPI
	buttonsConfig := configs.GetButtonsConfig()
	stage := configs.GetStageConfig()

	user, ok := users[callbackQuery.Message.Chat.ID]
	if !ok {
		if usr := database.GetUserData(callbackQuery.Message.Chat.ID); usr != nil {
			users[callbackQuery.Message.Chat.ID] = usr
			user = users[callbackQuery.Message.Chat.ID]
			user.Registered = true
		} else {
			users[callbackQuery.Message.Chat.ID] = new(types.User)
			user = users[callbackQuery.Message.Chat.ID]
			user.Registered = false
		}
	}

	_, ok = orders[callbackQuery.Message.Chat.ID]
	if !ok {
		orders[callbackQuery.Message.Chat.ID] = new(types.Order)
	}

	if !user.Registered {
		switch callbackQuery.Data {
		case buttonsConfig.SelectPhoneNumberInputMethod.TelegramApiInput.CallbackData:
			bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID-1))
			bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))
			bot.Send(sendPhoneNumber(callbackQuery.Message))
			return
		case buttonsConfig.SelectPhoneNumberInputMethod.ManualInput.CallbackData:
			bot.Send(enterPhoneNumberManualy(callbackQuery.Message))
			return
		case buttonsConfig.DataValidation.Correct.CallbackData:
			go clearChat(bot, callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID-1)
			bot.Send(createNewUser(callbackQuery.Message))
			return
		case buttonsConfig.DataValidation.Incorrect.CallbackData:
			bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID-1))
			bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))
			users[callbackQuery.Message.Chat.ID].Stage = stage.Registration.EnterUserNameAgain
			bot.Send(enterUserName(callbackQuery.Message.Chat.ID))
			return
		default:
			bot.Send(enterUserName(callbackQuery.Message.Chat.ID))
			return
		}
	}

	splitCallbackData := strings.Split(callbackQuery.Data, "*")

	switch splitCallbackData[0] {
	case buttonsConfig.MeinMenu.Menu.CallbackData:
		bot.Send(openMenu(callbackQuery.Message))
	case buttonsConfig.MeinMenu.Settings.CallbackData:
		bot.Send(openSettingsMenu(callbackQuery.Message))
	case buttonsConfig.Menu.MenuBack.CallbackData:
		bot.Send(openMainMenu(callbackQuery.Message))
	case buttonsConfig.Settings.SettingsBack.CallbackData:
		bot.Send(openMainMenu(callbackQuery.Message))
	case buttonsConfig.Menu.MakeOrderAddContactPerson.CallbackData:
		bot.Send(changeContactPersonName(callbackQuery.Message))
	case buttonsConfig.Menu.MakeOrderNewAddress.CallbackData:
		bot.Send(makeOrderChangeAddress(callbackQuery.Message))
	case buttonsConfig.Menu.CallMe.CallbackData:
		bot.Send(callMe(callbackQuery.Message))
	case buttonsConfig.Menu.CallMeSaveData.CallbackData:
		makeOrderSaveData(bot, callbackQuery.Message)
	case buttonsConfig.Menu.CallMeCancel.CallbackData:
		bot.Send(openMenu(callbackQuery.Message))
	case buttonsConfig.Menu.MakeOrder.CallbackData:
		bot.Send(makeOrderText(callbackQuery.Message))
	case buttonsConfig.Menu.MakeOrderCancel.CallbackData:
		bot.Send(openMenu(callbackQuery.Message))
	case buttonsConfig.Menu.MakeOrderSaveData.CallbackData:
		makeOrderSaveData(bot, callbackQuery.Message)
	case buttonsConfig.Menu.IssuesManage.CallbackData:
		bot.Send(issuesManage(callbackQuery.Message))
	case buttonsConfig.Menu.IssuesList.CallbackData:
		bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))
		issuesList(bot, callbackQuery.Message)
	case buttonsConfig.Menu.HideIssuesList.CallbackData:
		hideIssuesList(bot, callbackQuery.Message, splitCallbackData[1])
	case buttonsConfig.Menu.HideByMessageID.CallbackData:
		messageID, _ := strconv.Atoi(splitCallbackData[1])
		bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, messageID))
	case buttonsConfig.Menu.IssuesManageBack.CallbackData:
		bot.Send(openMenu(callbackQuery.Message))
	case buttonsConfig.Menu.SelectAssessment.CallbackData:
		selection, _ := strconv.Atoi(splitCallbackData[1])
		orders[callbackQuery.Message.Chat.ID].IssueID = splitCallbackData[2]
		bot.Send(selectAssessment(callbackQuery.Message, uint8(selection)))
	case buttonsConfig.Menu.SendAssessment.CallbackData:
		selection, _ := strconv.Atoi(splitCallbackData[1])
		bot.Send(sendAssessment(callbackQuery.Message, uint8(selection)))

	case buttonsConfig.Settings.ChangeData.CallbackData:
		bot.Send(openChangeDataMenu(callbackQuery.Message))
	case buttonsConfig.ChangeData.ChangeBack.CallbackData:
		bot.Send(openSettingsMenu(callbackQuery.Message))
	case buttonsConfig.Settings.DeleteAccount.CallbackData:
		bot.Send(beginDeleteAccount(callbackQuery.Message))
	case buttonsConfig.DeleteAccountValidation.No.CallbackData:
		bot.Send(openSettingsMenu(callbackQuery.Message))
	case buttonsConfig.DeleteAccountValidation.Yes.CallbackData:
		bot.Send(deleteAccount(callbackQuery.Message))
	case buttonsConfig.ChangeData.ChangeName.CallbackData:
		bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))
		bot.Send(beginChangeName(callbackQuery.Message))
	case buttonsConfig.ChangeData.ChangeOrganization.CallbackData:
		bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))
		bot.Send(beginChangeOrganization(callbackQuery.Message))
	case buttonsConfig.ChangeData.ChangeAddress.CallbackData:
		bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))
		bot.Send(beginChangeAddress(callbackQuery.Message))
	case buttonsConfig.ChangeData.ChangeEmail.CallbackData:
		bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))
		bot.Send(beginChangeEmail(callbackQuery.Message))
	case buttonsConfig.ChangeData.ChangePhoneNumber.CallbackData:
		bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))
		bot.Send(beginChangePhoneNumber(callbackQuery.Message))
	case buttonsConfig.SelectPhoneNumberInputMethod.TelegramApiInput.CallbackData:
		bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID-1))
		bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))
		bot.Send(sendPhoneNumber(callbackQuery.Message))
	case buttonsConfig.SelectPhoneNumberInputMethod.ManualInput.CallbackData:
		bot.Send(enterPhoneNumberManualy(callbackQuery.Message))
	case buttonsConfig.DataValidation.Correct.CallbackData:
		bot.Send(changePhoneNumber(callbackQuery.Message))
	case buttonsConfig.DataValidation.Incorrect.CallbackData:
		bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID-1))
		bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))
		users[callbackQuery.Message.Chat.ID].Stage = stage.Registration.EnterUserNameAgain
		bot.Send(enterUserName(callbackQuery.Message.Chat.ID))
	}

}
