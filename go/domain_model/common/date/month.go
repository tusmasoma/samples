package date

import "time"

type Month int64

const (
	MONTH_JANUARY Month = iota + 1
	MONTH_FEBRUARY
	MONTH_MARCH
	MONTH_APRIL
	MONTH_MAY
	MONTH_JUNE
	MONTH_JULY
	MONTH_AUGUST
	MONTH_SEPTEMBER
	MONTH_OCTOBER
	MONTH_NOVEMBER
	MONTH_DECEMBER
)

var (
	Month_name = map[int64]time.Month{
		1:  time.January,
		2:  time.February,
		3:  time.March,
		4:  time.April,
		5:  time.May,
		6:  time.June,
		7:  time.July,
		8:  time.August,
		9:  time.September,
		10: time.October,
		11: time.November,
		12: time.December,
	}
	Month_value = map[time.Month]int64{
		time.January:   1,
		time.February:  2,
		time.March:     3,
		time.April:     4,
		time.May:       5,
		time.June:      6,
		time.July:      7,
		time.August:    8,
		time.September: 9,
		time.October:   10,
		time.November:  11,
		time.December:  12,
	}
)

func (m Month) Value() int64 {
	return int64(m)
}

func (m Month) ToTimeMonth() time.Month {
	return Month_name[int64(m)]
}

func (m Month) ToInt() int64 {
	return int64(m)
}
func (m Month) ToString() string {
	return m.ToTimeMonth().String()
}

func (m Month) ToStringShort() string {
	return m.ToTimeMonth().String()[:3]
}

func (m Month) ToStringFull() string {
	return m.ToTimeMonth().String()
}

func MonthOf(month time.Month) Month {
	return Month(Month_value[month])
}
