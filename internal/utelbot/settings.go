package utelbot

import (
	"codeberg.org/kalsin/UtelBot/pkg/buttons"
	"codeberg.org/kalsin/UtelBot/pkg/configs"
	"codeberg.org/kalsin/UtelBot/pkg/database"
	"codeberg.org/kalsin/UtelBot/pkg/validators"
	tgbotapi "github.com/mazanur/telegram-bot-api/v6"
)

func openSettingsMenu(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	users[message.Chat.ID].Stage = stage.Settings.Menu
	msg.Text = messages.SelectAction
	msg.ReplyMarkup = buttons.Settings()
	return msg
}

func openChangeDataMenu(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	users[message.Chat.ID].Stage = stage.Settings.ChangeData
	msg.Text = messages.SelectAction
	msg.ReplyMarkup = buttons.ChangeData()
	return msg
}

func openMainMenu(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	users[message.Chat.ID].Stage = ""
	msg.Text = messages.SelectAction
	msg.ReplyMarkup = buttons.MainMenu()
	return msg
}

func beginDeleteAccount(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")
	stage := configs.GetStageConfig()

	users[message.Chat.ID].Stage = stage.Settings.BeginDeleteAccount
	msg.Text = messages.SettingsBeginDeleteAccount
	msg.ReplyMarkup = buttons.DeleteAccountValidation()
	return msg
}

func deleteAccount(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	err := database.DeleteUser(message.Chat.ID)
	if err != nil {
		users[message.Chat.ID].Stage = ""
		msg.Text = messages.SettingsDeleteAccountError
		msg.ReplyMarkup = buttons.MainMenu()
		return msg
	}

	users[message.Chat.ID].Stage = ""
	users[message.Chat.ID].Registered = false
	msg.Text = messages.SettingsDeleteAccount
	return msg
}

func beginChangeName(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	users[message.Chat.ID].Stage = stage.Settings.ChangeName
	msg.Text = messages.EnterUserName
	return msg
}

func changeName(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	users[message.Chat.ID].Name = message.Text
	users[message.Chat.ID].Stage = stage.Settings.ChangeData
	msg.ReplyMarkup = buttons.ChangeData()
	err := database.ChangeName(users[message.Chat.ID])
	if err != nil {
		msg.Text = messages.SettingsChangeNameError
		return msg
	}

	msg.Text = messages.SettingsChangeName
	return msg
}

func beginChangeOrganization(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	users[message.Chat.ID].Stage = stage.Settings.ChangeOrganization
	msg.Text = messages.RegistrationEnterOrganization
	return msg
}

func changeOrganization(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	users[message.Chat.ID].Organization = message.Text
	users[message.Chat.ID].Stage = stage.Settings.ChangeData
	msg.ReplyMarkup = buttons.ChangeData()
	err := database.ChangeOrganization(users[message.Chat.ID])
	if err != nil {
		msg.Text = messages.SettingsChangeOrganizationError
		return msg
	}

	msg.Text = messages.SettingsChangeOrganization
	return msg
}

func beginChangeAddress(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	users[message.Chat.ID].Stage = stage.Settings.ChangeAddress
	msg.Text = messages.RegistrationEnterAddressOrganization
	return msg
}

func changeAddress(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	users[message.Chat.ID].Address = message.Text
	users[message.Chat.ID].Stage = stage.Settings.ChangeData
	msg.ReplyMarkup = buttons.ChangeData()
	err := database.ChangeAddressOrganization(users[message.Chat.ID])
	if err != nil {
		msg.Text = messages.SettingsChangeAddressError
		return msg
	}

	msg.Text = messages.SettingsChangeAddress
	return msg
}

func beginChangeEmail(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	users[message.Chat.ID].Stage = stage.Settings.ChangeEmail
	msg.Text = messages.RegistrationEnterEmail
	return msg
}

func changeEmail(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	valid := validators.ValidateEmail(message.Text)
	if !valid {
		users[message.Chat.ID].Stage = stage.Settings.ChangeEmail
		msg.Text = messages.RegistrationEnterEmailError
		return msg
	}

	users[message.Chat.ID].Email = message.Text
	users[message.Chat.ID].Stage = stage.Settings.ChangeData
	msg.ReplyMarkup = buttons.ChangeData()
	err := database.ChangeEmail(users[message.Chat.ID])
	if err != nil {
		msg.Text = messages.SettingsChangeEmailError
		return msg
	}

	msg.Text = messages.SettingsChangeEmail
	return msg
}

func beginChangePhoneNumber(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	users[message.Chat.ID].Stage = stage.Settings.SelectPhoneNumberInputMethod
	msg.Text = messages.RegistrationSelectPhoneNumberInputMethod
	msg.ReplyMarkup = buttons.SelectPhoneNumberInputMethod()
	return msg
}

func changePhoneNumber(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	if users[message.Chat.ID].Stage == stage.Settings.EnterPhoneNumberTelegramApi {
		if message.Contact != nil {
			num, _ := validators.ValidatePhoneNumber(message.Contact.PhoneNumber)
			users[message.Chat.ID].PhoneNumber = num
		}
	}
	users[message.Chat.ID].Stage = stage.Settings.ChangeData
	msg.ReplyMarkup = buttons.ChangeData()
	err := database.ChangePhoneNumber(users[message.Chat.ID])
	if err != nil {
		msg.Text = messages.SettingsChangePhoneNumberError
		return msg
	}

	msg.Text = messages.SettingsChangePhoneNumber
	return msg
}
