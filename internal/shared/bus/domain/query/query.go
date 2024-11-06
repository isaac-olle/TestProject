package query

import (
	"TestProject/internal/shared/bus/domain"
	"encoding/json"
)

type IQuery interface {
	QueryType() string
	domain.IEvent
}

type BasicQuery struct {
	id string
}

func NewBasicQuery(id string) *BasicQuery {
	return &BasicQuery{id}
}

func (this *BasicQuery) QueryType() string {
	return "basicQuery"
}

func (this *BasicQuery) Serialize() ([]byte, error) {
	jsonQuery := struct {
		Id string `json:"id"`
	}{
		this.id,
	}
	return json.Marshal(jsonQuery)
}

func (this BasicQuery) Id() string {
	return this.id
}
