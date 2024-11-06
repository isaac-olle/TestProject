package contracts

type IDomainEntity interface {
	ToResponse(params ...any) any
}
