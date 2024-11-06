package vo

import "TestProject/internal/shared/domain/vo"

type UserSurname struct {
	vo.BasicValueObject[string]
}

func NewUserSurname(surname string) (*UserSurname, error) {
	if surname == "" {
		return nil, nil
	}
	// AQUETS METODES RETORNEN ERROR PEREQUE AQUI ES PODEN FER VALIDACIONS EXTRES.
	return &UserSurname{vo.NewBasicValueObject[string](surname)}, nil
}
