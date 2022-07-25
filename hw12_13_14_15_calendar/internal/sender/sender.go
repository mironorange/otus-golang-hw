package sender

import (
	"fmt"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/app"
)

type ICanSend interface {
	Send(n Notification) error
}

type Sender struct {
	logger app.Logger
}

type Notification struct {
	Title string
	Body  string
}

func New(l app.Logger) ICanSend {
	return &Sender{
		logger: l,
	}
}

func (s *Sender) Send(n Notification) error {
	s.logger.Info(fmt.Sprintf("send notification: %s", n.Title))
	return nil
}
