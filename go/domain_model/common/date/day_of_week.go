package date

import (
	"time"
)

type DayOfWeek int64

const (
	DAYOFWEEK_SUNDAY DayOfWeek = iota
	DAYOFWEEK_MONDAY
	DAYOFWEEK_TUESDAY
	DAYOFWEEK_WEDNESDAY
	DAYOFWEEK_THURSDAY
	DAYOFWEEK_FRIDAY
	DAYOFWEEK_SATURDAY
)

var (
	DayOfWeek_name = map[int64]time.Weekday{
		0: time.Sunday,
		1: time.Monday,
		2: time.Tuesday,
		3: time.Wednesday,
		4: time.Thursday,
		5: time.Friday,
		6: time.Saturday,
	}
	DayOfWeek_value = map[time.Weekday]int64{
		time.Sunday:    0,
		time.Monday:    1,
		time.Tuesday:   2,
		time.Wednesday: 3,
		time.Thursday:  4,
		time.Friday:    5,
		time.Saturday:  6,
	}
)

func (d DayOfWeek) ToTimeWeekday() time.Weekday {
	return DayOfWeek_name[int64(d)]
}

func (d DayOfWeek) ToInt() int64 {
	return int64(d)
}

func (d DayOfWeek) ToString() string {
	return d.ToTimeWeekday().String()
}

func (d DayOfWeek) ToStringShort() string {
	return d.ToTimeWeekday().String()[:3]
}

func (d DayOfWeek) ToStringFull() string {
	return d.ToTimeWeekday().String()
}

func DayOfWeekOf(weekday time.Weekday) DayOfWeek {
	return DayOfWeek(DayOfWeek_value[weekday])
}
