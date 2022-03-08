package db_handler

import (
	"fmt"
	"github.com/higuruchi/bulletin-board-sample.git/internal/message"
	"github.com/higuruchi/bulletin-board-sample.git/internal/db_handler/text_gateway"
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor/db_gateway"
)

type textHandler struct {
	database text_gateway.TextGateway
}

func New(database text_gateway.TextGateway) db_gateway.DB {
	return &textHandler {
		database: database,
	}
}

func (t *textHandler)GetAllMessage() ([]message.Message, error) {
	resMessage := []message.Message{}

	for i := 0; true; i++ {
		message_string, err := t.database.Get(i)
		if err != nil {
			break
		}

		message := message.New(message_string)
		resMessage = append(resMessage, message)
	}

	return resMessage, nil
}

func (t *textHandler)RecordMessage(m message.Message) error {
	err := t.database.Append(m.Message())
	if err != nil {
		return fmt.Errorf("Calling db_handler.text_handler.RecordMessage: %w", err)
	}

	return nil
}