package domain

import (
	"github.com/google/uuid"
)

// Amb un IdGenerator podriem estar controlant fer id AUTOINCREMENTALS o aixi
type IIdGenerator interface {
	Generate() uuid.UUID
}
