package db_handler

import (
	"fmt"
	"github.com/higuruchi/bulletin-board-sample.git/internal/message"
	"github.com/higuruchi/bulletin-board-sample.git/internal/db_handler/mysql_gateway"
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor/db_gateway"
)

type mySQLHandler struct {
	database mysql_gateway.MySQLHandler
}

func NewMySQL(database mysql_gateway.MySQLHandler) db_gateway.DB {
	return &mySQLHandler {
		database: database,
	}
}

func (m *mySQLHandler)GetAllMessage()(
	[]message.Message,
	error,
) {
	sql := `
	SELECT *
	FROM messages
	`
	rows, err := m.database.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("Calling db_handler.mysql_handler.GetAllMessage: %w", err)
	}

	messages := []message.Message{}

	for rows.Next() {
		var messageDetail string

		err = rows.Scan(&messageDetail)
		if err != nil {
			return nil, fmt.Errorf("Calling db_handler.mysql_handler.GetAllMessage: %w", err)
		}

		messages = append(messages, message.New(messageDetail))
	}

	return messages, nil
}

func (m *mySQLHandler)RecordMessage(
	message message.Message,
) error {
	sql := `
	INSERT INTO messages
	(message)
	VALUES (?)
	`
	_, err := m.database.Execute(sql, message.Message())
	if err != nil {
		return fmt.Errorf("Calling db_handler.mysql_handler.RecordMessage: %w", err)
	}

	return nil
}