package employee

import (
	"errors"
	"regexp"
)

const mailAddressPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

type MailAddress struct {
	Value string
}

func NewMailAddress(value string) (*MailAddress, error) {
	if value == "" {
		return nil, errors.New("mainl address is empty")
	}
	matched, err := regexp.MatchString(mailAddressPattern, value)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.New("invalid email address format")
	}
	return &MailAddress{Value: value}, nil
}

func (m *MailAddress) String() string {
	return m.Value
}
