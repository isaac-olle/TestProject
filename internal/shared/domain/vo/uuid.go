package vo

import (
	"github.com/google/uuid"
)

//No hi ha comprovacions perque les comprovacions es fan a traves dels validators de gin-gonic

type Uuid uuid.UUID

func NewUuid(uuid uuid.UUID) Uuid {
	return Uuid(uuid)
}

func (this Uuid) ToString() string {
	id := uuid.UUID(this)
	return id.String()
}

func (this Uuid) Value() uuid.UUID {
	return uuid.UUID(this)
}
