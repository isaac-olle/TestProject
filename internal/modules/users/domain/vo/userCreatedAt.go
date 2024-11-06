package vo

import (
	error2 "TestProject/internal/error"
	"TestProject/internal/shared/domain/vo"
	"fmt"
	"time"
)

type UserCreatedAt struct {
	vo.Date
}

func NewCreatedAt(t *time.Time) (*UserCreatedAt, error) {
	//AQUI ES PODEN FER VALIDACIONS EXTRES
	return &UserCreatedAt{vo.NewDate(t)}, nil
}

func NewCreatedAtFromString(createdAt string, format string) (*UserCreatedAt, error) {
	if createdAt == "" {
		return nil, nil
	}
	parsedTime, err := time.Parse(format, createdAt)
	if err != nil {
		return nil, error2.NewBadRequestHttpError(fmt.Sprintf("invalid date format: %s", err))
	}
	return NewCreatedAt(&parsedTime)
}
