package date

import (
	"fmt"
	"time"
)

// YearMonth は年月を表す構造体
// 例: 2025-04
type YearMonth struct {
	value time.Time
}

// NewYearMonthNow は現在の年月を返す
func NewYearMonthNow() *YearMonth {
	now := time.Now()
	return &YearMonth{value: time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)}
}

// NewYearMonth は年と月を指定して作成
func NewYearMonth(year *Year, month Month) *YearMonth {
	return &YearMonth{
		value: time.Date(year.Value(), time.Month(month.Value()), 1, 0, 0, 0, 0, time.Local),
	}
}

// NewYearMonthFromInts は int 値から生成
func NewYearMonthFromInts(year int, month int) *YearMonth {
	return &YearMonth{
		value: time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local),
	}
}

// NewYearMonthFromString は "2025-04" のような文字列から生成
func NewYearMonthFromString(ymStr string) (*YearMonth, error) {
	t, err := time.Parse("2006-01", ymStr)
	if err != nil {
		return nil, fmt.Errorf("月のフォーマットが誤っています: %w", err)
	}
	return &YearMonth{value: t}, nil
}

// Year を返す
func (ym *YearMonth) Year() *Year {
	return NewYear(ym.value.Year())
}

// Month を返す
func (ym *YearMonth) Month() Month {
	return Month(ym.value.Month())
}

// Start はその月の1日を返す
func (ym *YearMonth) Start() time.Time {
	return ym.value
}

// End はその月の末日を返す
func (ym *YearMonth) End() time.Time {
	return ym.value.AddDate(0, 1, -1)
}

// Days はその月の日付一覧を返す
func (ym *YearMonth) Days() []time.Time {
	start := ym.Start()
	end := ym.End()
	var days []time.Time
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}
	return days
}

// Before は前の月を返す
func (ym *YearMonth) Before() *YearMonth {
	return &YearMonth{value: ym.value.AddDate(0, -1, 0)}
}

// After は次の月を返す
func (ym *YearMonth) After() *YearMonth {
	return &YearMonth{value: ym.value.AddDate(0, 1, 0)}
}

// String は "2025-04" のような形式で年月を返す
func (ym *YearMonth) String() string {
	return ym.value.Format("2006-01")
}

// IsThisYear はその年月が今年であるかを判定
func (ym *YearMonth) IsThisYear() bool {
	now := time.Now()
	return ym.value.Year() == now.Year()
}

// Value は内部保持している time.Time（その月の1日）を返す
func (ym *YearMonth) Value() time.Time {
	return ym.value
}
