package error

type IErrorRepository interface {
	RecordError(id string, err error)
	GetError(id string) ([]byte, error)
}
