package db_gateway

import (
	"github.com/higuruchi/bulletin-board-sample.git/internal/message"
)

type DB interface {
	GetAllMessage() ([]message.Message, error)
	RecordMessage(message.Message) error
}