package vo

import "TestProject/internal/shared/domain/vo"

type UserName struct {
	vo.BasicValueObject[string]
}

func NewUserName(name string) (*UserName, error) {
	if name == "" {
		return nil, nil
	}
	// AQUETS METODES RETORNEN ERROR PEREQUE AQUI ES PODEN FER VALIDACIONS EXTRES.
	return &UserName{vo.NewBasicValueObject[string](name)}, nil
}
