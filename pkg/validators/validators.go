package validators

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

func ValidateEmail(email string) bool {
	temp, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	splitEmail := strings.Split(temp.Address, "@")
	if len(splitEmail) != 2 {
		return false
	}

	domail := splitEmail[1]
	splitDomain := strings.Split(domail, ".")

	return len(splitDomain) == 2
}

func ValidatePhoneNumber(number string) (string, error) {
	num, err := phonenumbers.Parse(number, "RU")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d%d", *num.CountryCode, *num.NationalNumber), nil
}
