package get

import "TestProject/internal/shared/bus/domain/query"

const (
	GetUserQueryName = "get_user"
)

type GetUserQuery struct {
	*query.BasicQuery
}

func NewGetUserQuery(id string) *GetUserQuery {
	return &GetUserQuery{query.NewBasicQuery(id)}
}

func (this *GetUserQuery) QueryType() string {
	return GetUserQueryName
}
