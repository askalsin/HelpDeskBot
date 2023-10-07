package buttons

import (
	"fmt"

	"codeberg.org/kalsin/UtelBot/pkg/configs"
	tgbotapi "github.com/mazanur/telegram-bot-api/v6"
)

func SelectPhoneNumberInputMethod() *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().SelectPhoneNumberInputMethod
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.ManualInput.Text, buttons.ManualInput.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.TelegramApiInput.Text, buttons.TelegramApiInput.CallbackData),
		),
	)
	return &keyboard
}

func SendPhoneNumber() *tgbotapi.ReplyKeyboardMarkup {
	buttons := configs.GetButtonsConfig().SelectPhoneNumberInputMethod.TelegramApiInput
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact(buttons.Text),
		),
	)
	return &keyboard
}

func DataValidation() *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().DataValidation
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.Correct.Text, buttons.Correct.CallbackData),
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.Incorrect.Text, buttons.Incorrect.CallbackData),
		),
	)
	return &keyboard
}

func MainMenu() *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().MeinMenu
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.Menu.Text, buttons.Menu.CallbackData),
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.Settings.Text, buttons.Settings.CallbackData),
		),
	)
	return &keyboard
}

func Menu() *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().Menu
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.InlineKeyboardButton{
				Text:   buttons.PriceList.Text,
				WebApp: &tgbotapi.WebAppInfo{URL: buttons.PriceList.CallbackData},
			},
		),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData(
		// 		buttons.IssuesManage.Text, buttons.IssuesManage.CallbackData),
		// ),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.IssuesList.Text, buttons.IssuesList.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.MakeOrder.Text, buttons.MakeOrder.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.CallMe.Text, buttons.CallMe.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.MenuBack.Text, buttons.MenuBack.CallbackData),
		),
	)
	return &keyboard
}

func MakeOrderCheckData() *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().Menu
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData(
		// 		buttons.MakeOrderAddPhoto.Text, buttons.MakeOrderAddPhoto.CallbackData),
		// ),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.MakeOrderAddContactPerson.Text, buttons.MakeOrderAddContactPerson.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.MakeOrderNewAddress.Text, buttons.MakeOrderNewAddress.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.MakeOrderSaveData.Text, buttons.MakeOrderSaveData.CallbackData),
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.MakeOrderCancel.Text, buttons.MakeOrderCancel.CallbackData),
		),
	)
	return &keyboard
}

func IssuesManage() *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().Menu
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.IssuesList.Text, buttons.IssuesList.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.DeleteIssues.Text, buttons.DeleteIssues.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.IssuesManageBack.Text, buttons.IssuesManageBack.CallbackData),
		),
	)
	return &keyboard
}

func CallMe() *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().Menu
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.CallMeSaveData.Text, buttons.CallMeSaveData.CallbackData),
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.CallMeCancel.Text, buttons.CallMeCancel.CallbackData),
		),
	)
	return &keyboard
}

func Settings() *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().Settings
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.DeleteAccount.Text, buttons.DeleteAccount.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.ChangeData.Text, buttons.ChangeData.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.SettingsBack.Text, buttons.SettingsBack.CallbackData),
		),
	)
	return &keyboard
}

func ChangeData() *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().ChangeData
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.ChangeName.Text, buttons.ChangeName.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.ChangeOrganization.Text, buttons.ChangeOrganization.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.ChangeAddress.Text, buttons.ChangeAddress.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.ChangePhoneNumber.Text, buttons.ChangePhoneNumber.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.ChangeEmail.Text, buttons.ChangeEmail.CallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.ChangeBack.Text, buttons.ChangeBack.CallbackData),
		),
	)
	return &keyboard
}

func DeleteAccountValidation() *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().DeleteAccountValidation
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.Yes.Text, buttons.Yes.CallbackData),
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.No.Text, buttons.No.CallbackData),
		),
	)
	return &keyboard
}

func HideIssuesList(len int) *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().Menu
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.HideIssuesList.Text,
				fmt.Sprintf("%s*%d", buttons.HideIssuesList.CallbackData, len)),
		),
	)
	return &keyboard
}

func HideByMessageID(messageID int) *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().Menu
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.HideByMessageID.Text,
				fmt.Sprintf("%s*%d", buttons.HideByMessageID.CallbackData, messageID)),
		),
	)
	return &keyboard
}

func Assessment(selection uint8, issueID string) *tgbotapi.InlineKeyboardMarkup {
	buttons := configs.GetButtonsConfig().Menu
	var buttonsArray []tgbotapi.InlineKeyboardButton

	for i := 0; i < 10; i++ {
		text := fmt.Sprintf("%d", i+1)
		callbackData := fmt.Sprintf("%s*%d*%s", buttons.SelectAssessment.CallbackData, i+1, issueID)
		btn := tgbotapi.NewInlineKeyboardButtonData(text, callbackData)
		buttonsArray = append(buttonsArray, btn)
	}

	var keyboard tgbotapi.InlineKeyboardMarkup

	if selection > 0 {
		buttonsArray[selection-1].Text = fmt.Sprintf("·%d·", selection)

		callbackData := fmt.Sprintf("%s*%d", buttons.SendAssessment.CallbackData, selection)
		keyboard = tgbotapi.NewInlineKeyboardMarkup(
			buttonsArray[:5],
			buttonsArray[5:],
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(buttons.SendAssessment.Text, callbackData),
			),
		)
	} else {
		keyboard = tgbotapi.NewInlineKeyboardMarkup(
			buttonsArray[:5],
			buttonsArray[5:],
		)
	}

	return &keyboard
}
