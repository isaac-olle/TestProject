package vo

import "TestProject/internal/shared/domain/vo"

type UserUsername struct {
	vo.BasicValueObject[string]
}

func NewUserUsername(username string) (*UserUsername, error) {
	if username == "" {
		return nil, nil
	}
	// AQUETS METODES RETORNEN ERROR PEREQUE AQUI ES PODEN FER VALIDACIONS EXTRES.
	return &UserUsername{vo.NewBasicValueObject[string](username)}, nil
}
