package vo

import "TestProject/internal/shared/domain/vo"

type UserPassword struct {
	vo.BasicValueObject[string]
}

func NewUserPassword(password string) (*UserPassword, error) {
	if password == "" {
		return nil, nil
	}
	// AQUETS METODES RETORNEN ERROR PEREQUE AQUI ES PODEN FER VALIDACIONS EXTRES.
	return &UserPassword{vo.NewBasicValueObject[string](password)}, nil
}
