package domain

type IEvent interface {
	Serialize() ([]byte, error)
	Id() string
}
