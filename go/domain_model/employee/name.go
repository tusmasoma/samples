package employee

import (
	"errors"
	"fmt"
)

const maxNameLength = 40

type Name struct {
	Value string
}

func NewName(value string) (*Name, error) {
	if value == "" {
		return nil, errors.New("name is empty")
	}
	if len(value) > maxNameLength {
		return nil, fmt.Errorf("name length exceeds %d characters", maxNameLength)
	}
	return &Name{Value: value}, nil
}

func (n *Name) String() string {
	return n.Value
}
