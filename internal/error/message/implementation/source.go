package implementation

type MessageSource struct {
	messages map[string]string
}

func NewMessageSource() *MessageSource {
	handler := MessageSource{}

	handler.messages = map[string]string{
		"api.request.body.incorrect":  "Incorrect request body",
		"api.request.param.incorrect": "Incorrect request param",
		"api.data.create.error":       "Error while creating data",
		"api.data.delete.error":       "Error while deleting data",
		"api.data.get.error":          "Error while getting data",
		"api.data.update.error":       "Error while updating data",
	}

	return &handler
}

func (m *MessageSource) GetMessage(error string) string {
	return m.messages[error]
}
