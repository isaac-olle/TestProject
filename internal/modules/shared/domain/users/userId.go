package users

import (
	cerror "TestProject/internal/error"
	"TestProject/internal/shared/domain/vo"
	"fmt"
	"github.com/google/uuid"
)

type UserId struct {
	vo.Uuid
}

func NewUserId(id uuid.UUID) (*UserId, error) {
	// AQUETS METODES RETORNEN ERROR PEREQUE AQUI ES PODEN FER VALIDACIONS EXTRES.
	return &UserId{vo.NewUuid(id)}, nil
}

func NewUserIdFromString(id string) (*UserId, error) {
	if id == "" {
		return nil, nil
	}
	idParsed, err := uuid.Parse(id)
	if err != nil {
		return nil, cerror.NewBadRequestHttpError(fmt.Sprintf("invalid userId: %s", err.Error()))
	}
	return NewUserId(idParsed)
}
