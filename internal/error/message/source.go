package message

type messageSource struct {
	messages map[string]string
}

func NewMessageSource(messages map[string]string) *messageSource {
	return &messageSource{messages: messages}
}

func (m *messageSource) GetMessage(error string) string {
	return m.messages[error]
}
