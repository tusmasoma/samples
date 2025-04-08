package employee

import (
	"strconv"
)

type EmployeeNumber struct {
	Value int64
}

func NewEmployeeNumber(value int64) (*EmployeeNumber, error) {
	return &EmployeeNumber{Value: value}, nil
}

func (e *EmployeeNumber) String() string {
	return strconv.Itoa(int(e.Value))
}
