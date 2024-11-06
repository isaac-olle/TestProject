package domain

import "TestProject/internal/shared/domain/contracts"

type ISqlModel[T contracts.IDomainEntity] interface {
	LoadFromDB(scan func(dest ...any) error) error
	ToDomain() (T, error)
}
