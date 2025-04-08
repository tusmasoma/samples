package date

import (
	"fmt"
	"strconv"
)

// Year は年を表す構造体
// 例: 2025
type Year struct {
	value int
}

func NewYear(year int) *Year {
	return &Year{value: year}
}

func NewYearFromString(yearStr string) (*Year, error) {
	y, err := strconv.Atoi(yearStr)
	if err != nil {
		return nil, fmt.Errorf("invalid year string: %s", yearStr)
	}
	return &Year{value: y}, nil
}

func (y *Year) Value() int {
	return y.value
}

func (y *Year) String() string {
	return strconv.Itoa(y.value)
}

func (y *Year) SameValue(other *Year) bool {
	return y.value == other.value
}
