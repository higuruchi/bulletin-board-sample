package database

import (
	"os"
	"fmt"
	"bufio"
	"errors"
	"github.com/higuruchi/bulletin-board-sample.git/internal/db_handler/text_gateway"
)

var (
	IndexNotFound = errors.New("Index Not Found")
)

type text struct {
	file string
}

func NewText(fileName string) (text_gateway.TextGateway, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Calling database.text.New: %w", err)
	}
	defer file.Close()


	return &text{
		file: fileName,
	}, nil
}

func (t *text)Append(data string) error {
	f, err := os.OpenFile(t.file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Calling database.text.Append: %w", err)
	}
	defer f.Close()

	if _, err :=f.Write([]byte(fmt.Sprintf("%s\n", data))); err != nil {
		return fmt.Errorf("Calling database.text.Append: %w", err)
	}

	return nil
}

func (t *text)Get(index int) (string, error) {
	f, err := os.Open(t.file)
	if err != nil {
		return "", fmt.Errorf("Calling database.text.Get: %w", err)
	}
	defer f.Close()

	line := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if index == line {
			return scanner.Text(), nil
		}
		line++
	}

	return "", IndexNotFound
}