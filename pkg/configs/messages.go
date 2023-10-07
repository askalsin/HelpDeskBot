package configs

import (
	"fmt"
	"sync"

	log "codeberg.org/kalsin/UtelBot/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

var (
	Messages           MessagesConfig
	onceMessagesConfig sync.Once
)

type MessagesConfig struct {
	WrongCommand                             string `yaml:"wrong_command"`
	SelectAction                             string `yaml:"select_action"`
	RegistrationEnterUserName                string `yaml:"registration_enter_user_name"`
	EnterUserName                            string `yaml:"enter_user_name"`
	RegistrationEnterOrganization            string `yaml:"registration_enter_organization"`
	RegistrationEnterAddressOrganization     string `yaml:"registration_enter_address_organization"`
	RegistrationSelectPhoneNumberInputMethod string `yaml:"registration_select_phone_number_input_method"`
	RegistrationEnterPhoneNumberTelegramApi  string `yaml:"registration_enter_phone_number_telegram_api"`
	RegistrationEnterPhoneNumberManualy      string `yaml:"registration_enter_phone_number_manualy"`
	RegistrationEnterPhoneNumberIncorrect    string `yaml:"registration_enter_phone_number_incorrect"`
	RegistrationEnterEmail                   string `yaml:"registration_enter_email"`
	RegistrationEnterEmailError              string `yaml:"registration_enter_email_error"`
	RegistrationCompleteRegistration         string `yaml:"registration_complete_registration"`
	RegistrationErrorCreateNewUser           string `yaml:"registration_error_create_new_user"`
	RegistrationСompletedSuccessfully        string `yaml:"registration_completed_successfully"`
	Initials                                 string `yaml:"initials"`
	Organization                             string `yaml:"organization"`
	Address                                  string `yaml:"address"`
	PhoneNumber                              string `yaml:"phone_number"`
	Email                                    string `yaml:"email"`
	Description                              string `yaml:"desctiprion"`
	Status                                   string `yaml:"status"`
	Assignee                                 string `yaml:"assignee"`
	ContactPerson                            string `yaml:"contact_person"`
	SettingsBeginDeleteAccount               string `yaml:"settings_begin_delete_account"`
	SettingsDeleteAccount                    string `yaml:"settings_delete_account"`
	SettingsDeleteAccountError               string `yaml:"settings_delete_account_error"`
	SettingsChangeName                       string `yaml:"settings_change_name"`
	SettingsChangeNameError                  string `yaml:"settings_change_name_error"`
	SettingsChangeOrganization               string `yaml:"settings_change_organization"`
	SettingsChangeOrganizationError          string `yaml:"settings_change_organization_error"`
	SettingsChangeAddress                    string `yaml:"settings_change_address"`
	SettingsChangeAddressError               string `yaml:"settings_change_address_error"`
	SettingsChangeEmail                      string `yaml:"settings_change_email"`
	SettingsChangeEmailError                 string `yaml:"settings_change_email_error"`
	SettingsChangePhoneNumber                string `yaml:"settings_change_phone_number"`
	SettingsChangePhoneNumberError           string `yaml:"settings_change_phone_number_error"`
	MenuMakeOrderText                        string `yaml:"menu_make_order_text"`
	MenuMakeOrderCheckData                   string `yaml:"menu_make_order_check_data"`
	OrderDataGroup                           string `yaml:"order_data_group"`
	MenuMakeOrderIssueSaved                  string `yaml:"menu_make_order_issue_saved"`
	MenuMakeOrderIssueSaveError              string `yaml:"menu_make_order_issue_save_error"`
	MenuMakeOrderChangeNameContactPerson     string `yaml:"menu_make_order_change_name_contact_person"`
	MenuMakeOrderChangeNumberContactPerson   string `yaml:"menu_make_order_change_number_contact_person"`
	CallMe                                   string `yaml:"call_me"`
	EmptyIssuesList                          string `yaml:"menu_empty_issues_list"`
	StartGroup                               string `yaml:"start_group"`
	ObserverChangeStatus                     string `yaml:"observer_change_status"`
	ObserverChangeAssignee                   string `yaml:"observer_change_assignee"`
	StartAssessment                          string `yaml:"start_assessment"`
	TextAssessment                           string `yaml:"text_assessment"`
	SendTextAssessment                       string `yaml:"send_text_assessment"`
	Assessment                               string `yaml:"assessment"`
	Comment                                  string `yaml:"comment"`
}

// Возвращает адрес структуры, в которой хранится конфигурация сообщений бота
func GetMessagesConfig() *MessagesConfig {
	onceMessagesConfig.Do(func() {
		Messages = MessagesConfig{}
		configFile := "configs/messages.yml"
		configPath := fmt.Sprintf("%s/%s", *GetCurrentWorkingDirectory(), configFile)

		if err := cleanenv.ReadConfig(configPath, &Messages); err != nil {
			log.Error.Fatalf("%s: error parse messages config!", err)
		}
	})

	return &Messages
}
