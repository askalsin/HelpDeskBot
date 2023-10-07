package utelbot

import (
	"fmt"
	
	"codeberg.org/kalsin/UtelBot/pkg/buttons"
	"codeberg.org/kalsin/UtelBot/pkg/configs"
	"codeberg.org/kalsin/UtelBot/pkg/database"
	"codeberg.org/kalsin/UtelBot/pkg/validators"
	tgbotapi "github.com/mazanur/telegram-bot-api/v6"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
)

func enterUserName(chatID int64) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(chatID, "")
	if users[chatID].Stage == stage.Registration.EnterUserNameAgain {
		msg.Text = messages.EnterUserName
	} else {
		msg.Text = messages.RegistrationEnterUserName
	}
	users[chatID].Stage = stage.Registration.EnterUserName

	msg.ParseMode = "Markdown"
	return msg
}

func enterOrganization(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.ReplyToMessageID = message.MessageID

	users[message.Chat.ID].Name = message.Text
	users[message.Chat.ID].Stage = stage.Registration.EnterOrganization
	msg.Text = messages.RegistrationEnterOrganization

	msg.ParseMode = "Markdown"
	return msg
}

func enterAddressOrganization(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.ReplyToMessageID = message.MessageID

	users[message.Chat.ID].Organization = message.Text
	users[message.Chat.ID].Stage = stage.Registration.EnterAddressOrganization
	msg.Text = messages.RegistrationEnterAddressOrganization

	msg.ParseMode = "Markdown"
	return msg
}

func selectPhoneNumberInputMethod(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.ReplyToMessageID = message.MessageID

	if users[message.Chat.ID].Stage == stage.Registration.EnterAddressOrganization {
		users[message.Chat.ID].Address = message.Text
		users[message.Chat.ID].Stage = stage.Registration.SelectPhoneNumberInputMethod
	} else if users[message.Chat.ID].Stage == stage.Settings.SelectPhoneNumberInputMethod {
		users[message.Chat.ID].Stage = stage.Settings.SelectPhoneNumberInputMethod
	}
	msg.Text = messages.RegistrationSelectPhoneNumberInputMethod
	msg.ReplyMarkup = buttons.SelectPhoneNumberInputMethod()

	msg.ParseMode = "Markdown"
	return msg
}

func sendPhoneNumber(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	if users[message.Chat.ID].Stage == stage.Registration.SelectPhoneNumberInputMethod {
		users[message.Chat.ID].Stage = stage.Registration.EnterPhoneNumberTelegramApi
	} else if users[message.Chat.ID].Stage == stage.Settings.SelectPhoneNumberInputMethod {
		users[message.Chat.ID].Stage = stage.Settings.EnterPhoneNumberTelegramApi
	}
	msg.Text = messages.RegistrationEnterPhoneNumberTelegramApi
	msg.ReplyMarkup = buttons.SendPhoneNumber()

	msg.ParseMode = "Markdown"
	return msg
}

func enterPhoneNumberManualy(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	if users[message.Chat.ID].Stage == stage.Registration.SelectPhoneNumberInputMethod {
		users[message.Chat.ID].Stage = stage.Registration.EnterPhoneNumberManualy
	} else if users[message.Chat.ID].Stage == stage.Settings.SelectPhoneNumberInputMethod {
		users[message.Chat.ID].Stage = stage.Settings.EnterPhoneNumberManualy
	}
	msg.Text = messages.RegistrationEnterPhoneNumberManualy

	msg.ParseMode = "Markdown"
	return msg
}

func phoneNumberValidating(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	num, err := validators.ValidatePhoneNumber(message.Text)
	if err != nil {
		msg.Text = messages.RegistrationEnterPhoneNumberIncorrect
		msg.ParseMode = "Markdown"
		return msg
	}
	users[message.Chat.ID].PhoneNumber = num

	if users[message.Chat.ID].Stage == stage.Registration.EnterPhoneNumberManualy {
		users[message.Chat.ID].Stage = stage.Registration.EnterEmail
		msg.Text = messages.RegistrationEnterEmail
	} else if users[message.Chat.ID].Stage == stage.Settings.EnterPhoneNumberManualy {
		msg = changePhoneNumber(message)
	}

	msg.ParseMode = "Markdown"
	return msg
}

func enterEmail(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.ReplyToMessageID = message.MessageID

	if message.Contact != nil {
		num, _ := validators.ValidatePhoneNumber(message.Contact.PhoneNumber)
		users[message.Chat.ID].PhoneNumber = num
	}

	users[message.Chat.ID].Stage = stage.Registration.EnterEmail
	msg.Text = messages.RegistrationEnterEmail

	msg.ParseMode = "Markdown"
	return msg
}

func completeRegistration(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.ParseMode = "Markdown"

	valid := validators.ValidateEmail(message.Text)
	if !valid {
		users[message.Chat.ID].Stage = stage.Registration.EnterEmail
		msg.Text = messages.RegistrationEnterEmailError
		return msg
	}
	users[message.Chat.ID].Email = message.Text
	users[message.Chat.ID].ChatID = message.Chat.ID
	users[message.Chat.ID].Stage = stage.Registration.CompleteRegistration
	msg.Text = fmt.Sprintf("%s\n%s%s\n%s%s\n%s%s\n%s%s\n%s%s",
		messages.RegistrationCompleteRegistration,
		messages.Initials, users[message.Chat.ID].Name,
		messages.Organization, users[message.Chat.ID].Organization,
		messages.Address, users[message.Chat.ID].Address,
		messages.PhoneNumber, users[message.Chat.ID].PhoneNumber,
		messages.Email, users[message.Chat.ID].Email)

	msg.ReplyMarkup = buttons.DataValidation()
	return msg
}

func createNewUser(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")
	msg.ParseMode = "Markdown"

	if err := database.NewUser(users[message.Chat.ID]); err != nil {
		log.Error.Printf("error create new user. User: %v| %s", users[message.Chat.ID], err)
		msg.Text = messages.RegistrationErrorCreateNewUser
		users[message.Chat.ID].Stage = stage.Registration.EnterUserName
		return msg
	}

	users[message.Chat.ID].Registered = true
	users[message.Chat.ID].Stage = ""
	msg.Text = fmt.Sprintf("%s\n%s", messages.RegistrationСompletedSuccessfully, messages.SelectAction)
	msg.ReplyMarkup = buttons.MainMenu()
	return msg
}
