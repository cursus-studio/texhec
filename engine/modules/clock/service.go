package clock

import "time"

type Service interface {
	GetDateFormat() DateFormat
	SetDateFormat(DateFormat)

	Now() time.Time
}

type DateFormat string

func NewDateFormat(date string) DateFormat { return DateFormat(date) }
func (format DateFormat) String() string   { return string(format) }

func (format DateFormat) Parse(date string) (time.Time, error) {
	return time.Parse(format.String(), date)
}
func (format DateFormat) Format(date time.Time) string { return date.Format(format.String()) }
