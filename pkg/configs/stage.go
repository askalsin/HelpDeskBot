package configs

import (
	"fmt"
	"sync"

	log "codeberg.org/kalsin/UtelBot/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

var (
	StageList       StageConfig
	onceStageConfig sync.Once
)

type StageConfig struct {
	Registration struct {
		EnterUserName                string `yaml:"enter_user_name"`
		EnterUserNameAgain           string `yaml:"enter_user_name_again"`
		EnterOrganization            string `yaml:"enter_organization"`
		EnterAddressOrganization     string `yaml:"enter_address_organization"`
		SelectPhoneNumberInputMethod string `yaml:"select_phone_number_input_method"`
		EnterPhoneNumberTelegramApi  string `yaml:"enter_phone_number_telegram_api"`
		EnterPhoneNumberManualy      string `yaml:"enter_phone_number_manualy"`
		EnterEmail                   string `yaml:"enter_email"`
		CompleteRegistration         string `yaml:"complete_registration"`
	} `yaml:"registration"`

	Menu struct {
		Open                               string `yaml:"open_menu"`
		MakeOrederText                     string `yaml:"make_order_text"`
		MakeOrderCheckData                 string `yaml:"make_order_check_data"`
		MakeOrderChangeNameContactPerson   string `yaml:"make_order_change_name_contact_person"`
		MakeOrderChangeNumberContactPerson string `yaml:"make_order_change_number_contact_person"`
		MakeOrderChangeNumberValidate      string `yaml:"make_order_change_number_validate"`
		MakeOrderNewAddress                string `yaml:"make_order_new_address"`
		CallMe                             string `yaml:"call_me"`
		SelectionAssissment                string `yaml:"selection_assissment"`
		SelectionText                      string `yaml:"selection_text"`
	} `yaml:"menu"`

	Settings struct {
		Menu                         string `yaml:"settings_menu"`
		ChangeData                   string `yaml:"settings_change_data"`
		BeginDeleteAccount           string `yaml:"settings_begin_delete_account"`
		ChangeName                   string `yaml:"change_name"`
		ChangeOrganization           string `yaml:"change_organization"`
		ChangeAddress                string `yaml:"change_address"`
		ChangeEmail                  string `yaml:"change_email"`
		ChangePhoneNumber            string `yaml:"change_phone_number"`
		SelectPhoneNumberInputMethod string `yaml:"select_phone_number_input_method"`
		EnterPhoneNumberTelegramApi  string `yaml:"enter_phone_number_telegram_api"`
		EnterPhoneNumberManualy      string `yaml:"enter_phone_number_manualy"`
	} `yaml:"settings"`
}

// Возвращает адрес структуры, в которой хранится конфигурация стадий бота
func GetStageConfig() *StageConfig {
	onceStageConfig.Do(func() {
		StageList = StageConfig{}
		configFile := "configs/stage.yml"
		configPath := fmt.Sprintf("%s/%s", *GetCurrentWorkingDirectory(), configFile)

		if err := cleanenv.ReadConfig(configPath, &StageList); err != nil {
			log.Error.Fatalf("%s: error parse stage config!", err)
		}
	})

	return &StageList
}
