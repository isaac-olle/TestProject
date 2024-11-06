package vo

import "time"

type CreatedAt struct {
	Date
}

func NewCreatedAt(t *time.Time) (*CreatedAt, error) {
	//AQUI ES PODEN FER VALIDACIONS EXTRES
	return &CreatedAt{NewDate(t)}, nil
}

func NewCreatedAtFromString(createdAt string, format string) (*CreatedAt, error) {
	if createdAt == "" {
		return nil, nil
	}
	parsedTime, err := time.Parse(format, createdAt)
	if err != nil {
		return nil, err
	}
	return NewCreatedAt(&parsedTime)
}
