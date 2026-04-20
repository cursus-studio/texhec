package clock

import (
	"time"

	"github.com/ogiusek/ioc/v2"
)

// interface

type DateFormat string

func NewDateFormat(date string) DateFormat { return DateFormat(date) }
func (format DateFormat) String() string   { return string(format) }

func (format DateFormat) Parse(date string) (time.Time, error) {
	return time.Parse(format.String(), date)
}
func (format DateFormat) Format(date time.Time) string { return date.Format(format.String()) }

// impl

type Clock interface {
	GetDateFormat() DateFormat
	SetDateFormat(DateFormat)

	Now() time.Time
}

type clock struct {
	format DateFormat
}

func (clock *clock) GetDateFormat() DateFormat  { return clock.format }
func (clock *clock) SetDateFormat(f DateFormat) { clock.format = f }

func (clock *clock) Now() time.Time {
	return time.Now()
}

// package

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Clock {
		return &clock{
			DateFormat(time.RFC3339Nano),
		}
	})
})
