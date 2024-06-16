package error

type Source interface {
	GetMessage(id string) string
}
