package message_getter

import (
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor/db_gateway"
)

type messageGetter struct {}

func New() db_gateway.DB {
	return &messageGetter{}
}

func (mg *messageGetter)GetAllMessage() error {
	return nil
}

func (mg *messageGetter)RecordMessage() error {
	return nil
}