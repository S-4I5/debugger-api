package message

type Source interface {
	GetMessage(id string) string
}
