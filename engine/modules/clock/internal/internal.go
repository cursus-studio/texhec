package internal

import (
	"engine/modules/clock"
	"time"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	format clock.DateFormat
}

func NewService(c ioc.Dic, format clock.DateFormat) clock.Service {
	return &service{
		format: format,
	}
}

func (clock *service) GetDateFormat() clock.DateFormat  { return clock.format }
func (clock *service) SetDateFormat(f clock.DateFormat) { clock.format = f }

func (clock *service) Now() time.Time {
	return time.Now()
}
