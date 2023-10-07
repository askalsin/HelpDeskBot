package observer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"codeberg.org/kalsin/UtelBot/pkg/buttons"
	"codeberg.org/kalsin/UtelBot/pkg/configs"
	"codeberg.org/kalsin/UtelBot/pkg/database"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
	"codeberg.org/kalsin/UtelBot/pkg/types"
	"codeberg.org/kalsin/UtelBot/pkg/yandextacker"
)

func Start() {
	log.Info.Println("Observer started!")

	for {
		issues := database.GetIssues()

		handleUpdates(issues)

		time.Sleep(10 * time.Minute)
	}
}

func handleUpdates(issues map[string]types.Issue) {
	conn := yandextacker.NewConnection(configs.GetTrackerConfig())
	messages := configs.GetMessagesConfig()

	for _, issue := range issues {
		data, err := conn.GetIssueDataByID(issue.IssueID)
		if err != nil {
			log.Error.Println(err)
			continue
		}

		if issue.Status != data.Status.Display {
			issue.Status = data.Status.Display
			_ = database.ChangeStatus(&issue)

			if issue.Status == "Закрыт" {
				issue.Status = data.Status.Display
				_ = database.ChangeStatus(&issue)

				text := fmt.Sprintf("%s\n%s", generateNotification(messages.ObserverChangeStatus, data), messages.StartAssessment)
				respData := sendMessage(issue.ChatID, text)
				keyboard := buttons.Assessment(0, issue.IssueID)
				byteKeyboard, _ := json.Marshal(keyboard)
				respData.editMessageReplyMarkup(issue.ChatID, byteKeyboard)
				continue
			}
			
			respData := sendMessage(issue.ChatID, generateNotification(messages.ObserverChangeStatus, data))
			keyboard := buttons.HideByMessageID(respData.Result.MessageID)
			byteKeyboard, _ := json.Marshal(keyboard)
			respData.editMessageReplyMarkup(issue.ChatID, byteKeyboard)
			continue
		}
		if issue.Assignee != data.Assignee.Name {
			issue.Assignee = data.Assignee.Name
			_ = database.ChangeAssignee(&issue)
			respData := sendMessage(issue.ChatID, generateNotification(messages.ObserverChangeAssignee, data))
			keyboard := buttons.HideByMessageID(respData.Result.MessageID)
			byteKeyboard, _ := json.Marshal(keyboard)
			respData.editMessageReplyMarkup(issue.ChatID, byteKeyboard)
			continue
		}
	}
}

func generateNotification(message string, data *yandextacker.ResponseIssue) string {
	messages := configs.GetMessagesConfig()

	var msg string
	if data.Assignee.Name == "" {
		msg = fmt.Sprintf("%s\n%s%s\n%s%s\n",
			message,
			messages.Description, data.Summary,
			messages.Status, data.Status.Display)
	} else {
		msg = fmt.Sprintf("%s\n%s%s\n%s%s\n%s%s",
			message,
			messages.Description, data.Summary,
			messages.Assignee, data.Assignee.Name,
			messages.Status, data.Status.Display)
	}
	return msg
}

type dataMessage struct {
	Result struct {
		MessageID int `json:"message_id"`
	} `json:"result"`
}

func sendMessage(chatID int64, text string) *dataMessage {
	resp, err := http.PostForm(
		fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", configs.GetToken()),
		url.Values{
			"chat_id":    {strconv.Itoa(int(chatID))},
			"text":       {text},
			"parse_mode": {"Markdown"}})

	if err != nil {
		log.Warning.Printf("Send message error -> %v", err)
	}

	data := dataMessage{}
	json.NewDecoder(resp.Body).Decode(&data)

	return &data
}

func (data *dataMessage) editMessageReplyMarkup(chatID int64, byteKeyboard []byte) {
	strMessageID := strconv.Itoa(data.Result.MessageID)
	// keyboard := buttons.HideByMessageID(strMessageID)
	// byteKeyboard, _ := json.Marshal(keyboard)

	_, err := http.PostForm(
		fmt.Sprintf("https://api.telegram.org/bot%v/editMessageReplyMarkup", configs.GetToken()),
		url.Values{
			"chat_id":      {strconv.Itoa(int(chatID))},
			"message_id":   {strMessageID},
			"reply_markup": {string(byteKeyboard)}})

	if err != nil {
		log.Warning.Printf("Edit message error -> %v", err)
	}
}

// func deleteMessage(userID int64, messageID int) {
// 	strMessageID := strconv.Itoa(messageID)
// 	_, err := http.PostForm(
// 		fmt.Sprintf("https://api.telegram.org/bot%v/deleteMessage", configs.GetToken()),
// 		url.Values{"chat_id": {strconv.Itoa(int(userID))}, "message_id": {strMessageID}})
// 	if err != nil {
// 		log.Warning.Printf("Delete message error -> %v", err)
// 	}
// }
