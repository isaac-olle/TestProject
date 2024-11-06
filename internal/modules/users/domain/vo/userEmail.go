package vo

import "TestProject/internal/shared/domain/vo"

type UserEmail struct {
	vo.BasicValueObject[string]
}

func NewUserEmail(email string) (*UserEmail, error) {
	if email == "" {
		return nil, nil
	}
	// AQUETS METODES RETORNEN ERROR PEREQUE AQUI ES PODEN FER VALIDACIONS EXTRES.
	return &UserEmail{vo.NewBasicValueObject[string](email)}, nil
}
