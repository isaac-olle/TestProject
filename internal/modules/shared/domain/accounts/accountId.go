package accounts

import (
	cerror "TestProject/internal/error"
	"TestProject/internal/shared/domain/vo"
	"github.com/google/uuid"
)

type AccountId struct {
	vo.Uuid
}

func NewAccountId(id uuid.UUID) (*AccountId, error) {
	// AQUETS METODES RETORNEN ERROR PEREQUE AQUI ES PODEN FER VALIDACIONS EXTRES.
	return &AccountId{vo.NewUuid(id)}, nil
}

func NewAccountIdFromString(id string) (*AccountId, error) {
	if id == "" {
		return nil, nil
	}
	idParsed, err := uuid.Parse(id)
	if err != nil {
		return nil, cerror.NewHttpError(404, err.Error())
	}
	return NewAccountId(idParsed)
}
