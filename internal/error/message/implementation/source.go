package implementation

type MessageSource struct {
	messages map[string]string
}

func NewMessageSource() *MessageSource {
	handler := MessageSource{}

	handler.messages = map[string]string{
		"id": "Id",
	}

	return &handler
}

func (m *MessageSource) GetMessage(error string) string {
	return m.messages[error]
}
