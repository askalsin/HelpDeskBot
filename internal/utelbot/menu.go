package utelbot

import (
	"fmt"
	"strconv"
	"strings"

	"codeberg.org/kalsin/UtelBot/pkg/buttons"
	"codeberg.org/kalsin/UtelBot/pkg/configs"
	"codeberg.org/kalsin/UtelBot/pkg/database"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
	"codeberg.org/kalsin/UtelBot/pkg/validators"
	"codeberg.org/kalsin/UtelBot/pkg/yandextacker"
	tgbotapi "github.com/mazanur/telegram-bot-api/v6"
)

func openMenu(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	users[message.Chat.ID].Stage = stage.Menu.Open
	msg.Text = messages.SelectAction
	msg.ReplyMarkup = buttons.Menu()

	msg.ParseMode = "Markdown"
	return msg
}

func makeOrderText(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	users[message.Chat.ID].Stage = stage.Menu.MakeOrederText
	msg.Text = messages.MenuMakeOrderText

	msg.ParseMode = "Markdown"
	return msg
}

func makeOrderCheckData(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	switch users[message.Chat.ID].Stage {
	case stage.Menu.MakeOrederText:
		orders[message.Chat.ID].ProblemDescription = message.Text
		orders[message.Chat.ID].ContactPerson = users[message.Chat.ID].Name
		orders[message.Chat.ID].PhoneNumberContactPerson = users[message.Chat.ID].PhoneNumber
		orders[message.Chat.ID].Organization = users[message.Chat.ID].Organization
		orders[message.Chat.ID].Address = users[message.Chat.ID].Address
		orders[message.Chat.ID].Email = users[message.Chat.ID].Email
	case stage.Menu.MakeOrderNewAddress:
		orders[message.Chat.ID].Address = message.Text
	}

	users[message.Chat.ID].Stage = stage.Menu.MakeOrderCheckData
	msg.Text = fmt.Sprintf("%s\n%s%s\n%s%s\n%s%s\n%s%s\n%s%s\n%s%s",
		messages.MenuMakeOrderCheckData,
		messages.Description, orders[message.Chat.ID].ProblemDescription,
		messages.Initials, orders[message.Chat.ID].ContactPerson,
		messages.Organization, orders[message.Chat.ID].Organization,
		messages.Address, orders[message.Chat.ID].Address,
		messages.PhoneNumber, orders[message.Chat.ID].PhoneNumberContactPerson,
		messages.Email, orders[message.Chat.ID].Email)
	msg.ReplyMarkup = buttons.MakeOrderCheckData()

	msg.ParseMode = "Markdown"
	return msg
}

func makeOrderSaveData(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	conn := yandextacker.NewConnection(configs.GetTrackerConfig())

	issues[message.Chat.ID] = yandextacker.NewRequestIssue()

	issues[message.Chat.ID].Summary = fmt.Sprintf("%s: %s",
		orders[message.Chat.ID].Organization, orders[message.Chat.ID].ProblemDescription)

	issues[message.Chat.ID].Description = fmt.Sprintf("%s%s\n%s%s\n%s%s\n%s%s",
		messages.Initials, orders[message.Chat.ID].ContactPerson,
		messages.Address, orders[message.Chat.ID].Address,
		messages.PhoneNumber, orders[message.Chat.ID].PhoneNumberContactPerson,
		messages.Email, orders[message.Chat.ID].Email)

	msg.Text = messages.MenuMakeOrderIssueSaved

	issueID, err := conn.CreateNewIssue(issues[message.Chat.ID])
	if err != nil {
		log.Error.Println(err)
		msg.Text = messages.MenuMakeOrderIssueSaveError
		msg.ReplyMarkup = buttons.Menu()
		bot.Send(msg)
		return
	}

	orders[message.Chat.ID].IssueID = issueID
	orders[message.Chat.ID].ClientChatID = message.Chat.ID
	orders[message.Chat.ID].Status = "Открыт"

	err = database.NewIssue(orders[message.Chat.ID])
	if err != nil {
		log.Error.Println(err)
		msg.Text = messages.MenuMakeOrderIssueSaveError
		msg.ReplyMarkup = buttons.Menu()
		bot.Send(msg)
		return
	}

	issues[message.Chat.ID] = yandextacker.NewRequestIssue()

	groupsID := database.GetGroupsChatID()
	for _, chatID := range groupsID {
		msgGroup := tgbotapi.NewMessage(chatID, "")
		msgGroup.Text = fmt.Sprintf("%s\n%s%s\n%s%s\n%s%s\n%s%s\n%s%s\n%s%s",
			messages.OrderDataGroup,
			messages.Description, orders[message.Chat.ID].ProblemDescription,
			messages.Initials, orders[message.Chat.ID].ContactPerson,
			messages.Organization, orders[message.Chat.ID].Organization,
			messages.Address, orders[message.Chat.ID].Address,
			messages.PhoneNumber, orders[message.Chat.ID].PhoneNumberContactPerson,
			messages.Email, orders[message.Chat.ID].Email)

		_, err := bot.Send(msgGroup)
		if err != nil {
			log.Error.Printf("error send to the -> (%d) Error:%v\n", chatID, err)
		}
	}

	users[message.Chat.ID].Stage = stage.Menu.Open
	msg.ReplyMarkup = buttons.Menu()

	msg.ParseMode = "Markdown"

	go clearChat(bot, message.Chat.ID, message.MessageID-1)
	bot.Send(msg)
}

func changeContactPersonName(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	users[message.Chat.ID].Stage = stage.Menu.MakeOrderChangeNameContactPerson
	msg.Text = messages.MenuMakeOrderChangeNameContactPerson

	msg.ParseMode = "Markdown"
	return msg
}

func changeContactPersonPhoneNumber(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.ReplyToMessageID = message.MessageID

	orders[message.Chat.ID].ContactPerson = message.Text
	users[message.Chat.ID].Stage = stage.Menu.MakeOrderChangeNumberContactPerson
	msg.Text = messages.MenuMakeOrderChangeNumberContactPerson

	msg.ParseMode = "Markdown"
	return msg
}

func changeContactPersonValidate(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	num, err := validators.ValidatePhoneNumber(message.Text)
	if err != nil {
		msg.Text = messages.RegistrationEnterPhoneNumberIncorrect
		return msg
	}
	orders[message.Chat.ID].PhoneNumberContactPerson = num

	users[message.Chat.ID].Stage = stage.Menu.MakeOrderChangeNumberValidate

	return makeOrderCheckData(message)
}

func makeOrderChangeAddress(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	users[message.Chat.ID].Stage = stage.Menu.MakeOrderNewAddress
	msg.Text = messages.RegistrationEnterAddressOrganization

	msg.ParseMode = "Markdown"
	return msg
}

func callMe(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	orders[message.Chat.ID].ProblemDescription = messages.CallMe
	orders[message.Chat.ID].ContactPerson = users[message.Chat.ID].Name
	orders[message.Chat.ID].PhoneNumberContactPerson = users[message.Chat.ID].PhoneNumber
	orders[message.Chat.ID].Organization = users[message.Chat.ID].Organization
	orders[message.Chat.ID].Address = users[message.Chat.ID].Address
	orders[message.Chat.ID].Email = users[message.Chat.ID].Email

	users[message.Chat.ID].Stage = stage.Menu.CallMe
	msg.Text = fmt.Sprintf("%s\n%s%s\n%s%s\n%s%s\n%s%s\n%s%s\n%s%s",
		messages.MenuMakeOrderCheckData,
		messages.Description, orders[message.Chat.ID].ProblemDescription,
		messages.Initials, orders[message.Chat.ID].ContactPerson,
		messages.Organization, orders[message.Chat.ID].Organization,
		messages.Address, orders[message.Chat.ID].Address,
		messages.PhoneNumber, orders[message.Chat.ID].PhoneNumberContactPerson,
		messages.Email, orders[message.Chat.ID].Email)

	msg.ReplyMarkup = buttons.CallMe()

	msg.ParseMode = "Markdown"
	return msg
}

func issuesManage(message *tgbotapi.Message) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	msg.Text = messages.SelectAction
	msg.ReplyMarkup = buttons.IssuesManage()

	msg.ParseMode = "Markdown"
	return msg
}

func issuesList(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	messages := configs.GetMessagesConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	conn := yandextacker.NewConnection(configs.GetTrackerConfig())

	issuesID, err := database.GetIssuesID(message.Chat.ID)
	if err != nil {
		log.Error.Println(err)
	}

	length := len(issuesID)

	if length == 0 {
		msg.Text = messages.EmptyIssuesList
		msg.ReplyMarkup = buttons.MainMenu()
		msg.ParseMode = "Markdown"
		bot.Send(msg)
		return
	}

	for i, val := range issuesID {
		issue, err := conn.GetIssueDataByID(val)
		if err != nil {
			log.Error.Println(err)
		}

		if issue.Assignee.Name == "" {
			msg.Text = fmt.Sprintf("%s%s\n%s%s\n",
				messages.Description, issue.Summary,
				messages.Status, issue.Status.Display)
		} else {
			msg.Text = fmt.Sprintf("%s%s\n%s%s\n%s%s",
				messages.Description, issue.Summary,
				messages.Assignee, issue.Assignee.Name,
				messages.Status, issue.Status.Display)
		}

		if i == length-1 {
			msg.ReplyMarkup = buttons.HideIssuesList(length)
		}

		msg.ParseMode = "Markdown"
		bot.Send(msg)
	}
}

func hideIssuesList(bot *tgbotapi.BotAPI, message *tgbotapi.Message, lengthString string) {
	messages := configs.GetMessagesConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	length, err := strconv.Atoi(lengthString)
	if err != nil {
		log.Error.Println(err)
	}

	for i := 0; i < length; i++ {
		go bot.Send(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-i))
	}

	msg.Text = messages.SelectAction
	msg.ReplyMarkup = buttons.MainMenu()
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}

func selectAssessment(message *tgbotapi.Message, selection uint8) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")

	users[message.Chat.ID].Stage = stage.Menu.SelectionAssissment
	msg.Text = messages.SelectAction
	msg.ReplyMarkup = buttons.Assessment(selection, "")

	msg.ParseMode = "Markdown"
	return msg
}

func sendAssessment(message *tgbotapi.Message, selection uint8) tgbotapi.EditMessageTextConfig {
	messages := configs.GetMessagesConfig()
	stage := configs.GetStageConfig()
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")
	msg.ParseMode = "Markdown"

	if selection < 8 {
		users[message.Chat.ID].Stage = fmt.Sprintf("%s*%d", stage.Menu.SelectionText, selection)
		msg.Text = messages.TextAssessment
		return msg
	}

	_ = database.DeleteIssue(orders[message.Chat.ID].IssueID)
	users[message.Chat.ID].Stage = ""
	msg.Text = messages.SelectAction
	msg.ReplyMarkup = buttons.MainMenu()

	return msg
}

func sendTextAssessment(message *tgbotapi.Message) tgbotapi.MessageConfig {
	messages := configs.GetMessagesConfig()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	conn := yandextacker.NewConnection(configs.GetTrackerConfig())

	splitStage := strings.Split(users[message.Chat.ID].Stage, "*")

	text := fmt.Sprintf("%s: %s\n%s: %s",
		messages.Assessment, splitStage[1],
		messages.Comment, message.Text,
	)

	err := conn.AddComment(orders[message.Chat.ID].IssueID, text)
	if err != nil {
		log.Error.Println(err)
	}

	_ = database.DeleteIssue(orders[message.Chat.ID].IssueID)

	msg.Text = messages.SendTextAssessment
	msg.ReplyMarkup = buttons.MainMenu()
	msg.ParseMode = "Markdown"

	return msg
}
