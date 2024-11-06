package command

import (
	"TestProject/internal/shared/bus/domain"
	"encoding/json"
)

type ICommand interface {
	CommandType() string
	domain.IEvent
}

func NewAbstractCommand(id string) *AbstractCommand {
	return &AbstractCommand{ID: id}
}

type AbstractCommand struct {
	ID string `json:"id,omitempty"`
}

func (this *AbstractCommand) Serialize() ([]byte, error) {
	return json.Marshal(this)
}

func (this *AbstractCommand) Id() string {
	return this.ID
}

func (this *AbstractCommand) SetId(id string) {
	this.ID = id
}
