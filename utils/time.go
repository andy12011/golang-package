package utils

import "time"

const (
	TIME_FORMAT_ONE = "2006-01-02 15:04:05" //YYYY-MM-DD HH:MM:SS
)

type timeFactory struct {
	Time time.Time
}

func NewTimeFactory() *timeFactory {
	return &timeFactory{
		Time: time.Now(),
	}
}

func (t *timeFactory) GetTime() time.Time {
	return t.Time
}

func (t *timeFactory) SetTime(newTime time.Time) *timeFactory {
	t.Time = newTime
	return t
}

func (t *timeFactory) ToString() string {
	return t.Time.Format(TIME_FORMAT_ONE)
}

func (t *timeFactory) Format(format string) string {
	return t.Time.Format(format)
}

func (t *timeFactory) Immutable() *timeFactory {
	nt := NewTimeFactory()
	nt.Time = t.Time

	return nt
}

func (t *timeFactory) AddHours(hours int) *timeFactory {
	t.Time = t.Time.Add(time.Duration(hours) * time.Hour)

	return t
}

func (t *timeFactory) SubHours(hours int) *timeFactory {
	t.Time = t.Time.Add(-time.Duration(hours) * time.Hour)

	return t
}

func (t *timeFactory) AddMinutes(minutes int) *timeFactory {
	t.Time = t.Time.Add(time.Duration(minutes) * time.Minute)
	return t
}

func (t *timeFactory) SubMinutes(minutes int) *timeFactory {
	t.Time = t.Time.Add(-time.Duration(minutes) * time.Minute)
	return t
}

func (t *timeFactory) AddSeconds(seconds int) *timeFactory {
	t.Time = t.Time.Add(time.Duration(seconds) * time.Second)
	return t
}

func (t *timeFactory) SubSeconds(seconds int) *timeFactory {
	t.Time = t.Time.Add(-time.Duration(seconds) * time.Second)
	return t
}

func (t *timeFactory) AddYears(years int) *timeFactory {
	t.Time = t.Time.AddDate(years, 0, 0)
	return t
}

func (t *timeFactory) SubYears(years int) *timeFactory {
	t.Time = t.Time.AddDate(-years, 0, 0)
	return t
}

func (t *timeFactory) AddMonths(months int) *timeFactory {
	t.Time = t.Time.AddDate(0, months, 0)
	return t
}

func (t *timeFactory) SubMonths(months int) *timeFactory {
	t.Time = t.Time.AddDate(0, -months, 0)
	return t
}

func (t *timeFactory) AddDays(days int) *timeFactory {
	t.Time = t.Time.AddDate(0, 0, days)
	return t
}

func (t *timeFactory) SubDays(days int) *timeFactory {
	t.Time = t.Time.AddDate(0, 0, -days)
	return t
}
