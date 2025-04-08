package employee

import (
	"errors"
	"regexp"
)

const (
	minPhoneNumberLength = 8
	maxPhoneNumberLength = 13
	phoneNumberPattern   = `^[0-9]{2,4}-[0-9]{2,4}-[0-9]{2,4}$`
)

type PhoneNumber struct {
	Value string
}

func NewPhoneNumber(value string) (*PhoneNumber, error) {
	if value == "" {
		return nil, errors.New("phone number is empty")
	}
	matched, err := regexp.MatchString(phoneNumberPattern, value)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.New("invalid phone number format")
	}
	if len(value) < minPhoneNumberLength || len(value) > maxPhoneNumberLength {
		return nil, errors.New("phone number length is invalid")
	}
	return &PhoneNumber{Value: value}, nil
}

func (p *PhoneNumber) String() string {
	return p.Value
}
