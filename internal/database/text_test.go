package database

import "testing"

func TestAppend(t *testing.T) {
	text, err := NewText("test.txt")
	if err != nil {
		panic(err)
	}

	if err := text.Append("append message"); err != nil {
		panic(err)
	}
}

func TestGet(t *testing.T) {
	text, err := NewText("test.txt")
	if err != nil {
		panic(err)
	}

	if m, err := text.Get(0); err != nil {
		panic(err)
	} else if m != "first message" {
		t.Errorf("want: %v got: %v", "first message", m)
	}
}