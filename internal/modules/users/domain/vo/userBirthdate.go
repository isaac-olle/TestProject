package vo

import (
	error2 "TestProject/internal/error"
	"TestProject/internal/shared/domain/vo"
	"fmt"
	"time"
)

type UserBirthdate struct {
	vo.Date
}

func NewUserBirthdate(birthdate *time.Time) (*UserBirthdate, error) {
	// AQUETS METODES RETORNEN ERROR PEREQUE AQUI ES PODEN FER VALIDACIONS EXTRES.
	return &UserBirthdate{vo.NewDate(birthdate)}, nil
}

func NewUserBirthdateFromString(birthdate, format string) (*UserBirthdate, error) {
	if birthdate == "" {
		return nil, nil
	}
	parsedTime, err := time.Parse(format, birthdate)
	if err != nil {
		return nil, error2.NewBadRequestHttpError(fmt.Sprintf("invalid birthdate: %s", err.Error()))
	}
	return NewUserBirthdate(&parsedTime)
}
