package infrastructure

import (
	"github.com/google/uuid"
)

type UuidGenerator struct {
}

func NewUuidGenerator() *UuidGenerator {
	return &UuidGenerator{}
}

func (this *UuidGenerator) Generate() uuid.UUID {
	return uuid.New()
}
