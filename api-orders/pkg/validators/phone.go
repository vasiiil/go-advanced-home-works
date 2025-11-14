package validators

import (
	"regexp"
	"strings"
)

func Phone(phone string) (string, error) {
	if phone == "" {
		return "", ErrEmptyPhone
	}
	matched, err := regexp.MatchString("^(7|8)?9\\d{9}$", phone)
	if err != nil {
		return "", err
	}
	if !matched {
		return "", ErrInvalidPhoneFormat
	}

	if !strings.HasPrefix(phone, "7") {
		phoneRunes := []rune(phone)
		if len(phoneRunes) == 10 {
			phoneRunes = append([]rune{'7'}, phoneRunes...)
		} else {
			phoneRunes[0] = '7'
		}

		phone = string(phoneRunes)
	}

	return phone, nil
}
