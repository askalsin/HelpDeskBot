package configs

import (
	"fmt"
	"sync"

	log "codeberg.org/kalsin/UtelBot/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

var (
	Buttons           ButtonsConfig
	onceButtonsConfig sync.Once
)

type ButtonsConfig struct {
	SelectPhoneNumberInputMethod struct {
		ManualInput struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"manual_input"`
		TelegramApiInput struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"telegram_api_input"`
	} `yaml:"select_phone_number_input_method"`

	DataValidation struct {
		Correct struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"correct"`
		Incorrect struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"incorrect"`
	} `yaml:"data_validation"`

	MeinMenu struct {
		Menu struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"menu"`
		Settings struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"settings"`
	} `yaml:"main_menu"`

	Menu struct {
		PriceList struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"price_list"`

		MakeOrder struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"make_order"`
		MakeOrderAddContactPerson struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"make_order_add_contact_person"`
		MakeOrderSaveData struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"make_order_save_data"`
		MakeOrderNewAddress struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"make_order_new_address"`
		MakeOrderCancel struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"make_order_cancel"`

		CallMe struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"call_me"`
		CallMeSaveData struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"call_me_save_data"`
		CallMeCancel struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"call_me_cancel"`

		IssuesManage struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"issues_manage"`
		IssuesList struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"issues_list"`
		DeleteIssues struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"delete_issues"`
		IssuesManageBack struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"issues_manage_back"`
		HideIssuesList struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"hide_issues_list"`
		HideByMessageID struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"hide_by_message_id"`
		SelectAssessment struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"select_assessment"`
		SendAssessment struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"send_assessment"`

		MenuBack struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"menu_back"`
	} `yaml:"menu"`

	Settings struct {
		DeleteAccount struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"delete_account"`
		ChangeData struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"change_data"`
		SettingsBack struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"settings_back"`
	} `yaml:"settings"`

	ChangeData struct {
		ChangeName struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"change_name"`
		ChangeOrganization struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"change_organization"`
		ChangeAddress struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"change_address"`
		ChangePhoneNumber struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"change_phone_number"`
		ChangeEmail struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"change_email"`
		ChangeBack struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"change_back"`
	} `yaml:"change_data"`

	DeleteAccountValidation struct {
		Yes struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"yes"`
		No struct {
			Text         string `yaml:"text"`
			CallbackData string `yaml:"callback_data"`
		} `yaml:"no"`
	} `yaml:"delete_account_validation"`
}

// Возвращает адрес структуры, в которой хранится конфигурация кнопок бота
func GetButtonsConfig() *ButtonsConfig {
	onceButtonsConfig.Do(func() {
		Buttons = ButtonsConfig{}
		configFile := "configs/buttons.yml"
		configPath := fmt.Sprintf("%s/%s", *GetCurrentWorkingDirectory(), configFile)

		if err := cleanenv.ReadConfig(configPath, &Buttons); err != nil {
			log.Error.Fatalf("%s: error parse buttons config!", err)
		}
	})

	return &Buttons
}
