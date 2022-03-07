package message

type Message interface {
	Message() string
}

type message struct {
	detail string
}

func New(detail string) Message {
	return &message {
		detail: detail,
	}
}

func (m *message)Message() string {
	return m.detail
}