package entities

import (
	error2 "TestProject/internal/error"
	"TestProject/internal/modules/users/domain/dtos"
	"TestProject/internal/modules/users/domain/entities"
	"database/sql"
	"errors"
)

const (
	timeFormat = "2006-01-02T15:04:05Z"
)

type UserSQLDatabaseTable struct {
	Id        string
	Name      string
	Surname   string
	Birthdate string
	Email     string
	CreatedAt string
	UpdatedAt string
	IsDeleted bool
	//Username  string
	//Password  string
}

func NewEmptyUserMySQLDatabaseTable() *UserSQLDatabaseTable {
	return &UserSQLDatabaseTable{}
}

func (this *UserSQLDatabaseTable) ToDomain() (*entities.User, error) {
	return entities.NewUserFromUnvaluedObjects(dtos.UserParams{
		Id:                  this.Id,
		Name:                this.Name,
		Surname:             this.Surname,
		Email:               this.Email,
		Birthdate:           this.Birthdate,
		BirthdateTimeFormat: timeFormat,
		CreatedAt:           this.CreatedAt,
		CreatedAtTimeFormat: timeFormat,
		//Username:            this.Username,
		//Password:            this.Password,
	})
}

func (this *UserSQLDatabaseTable) LoadFromDB(scan func(dest ...any) error) error {
	err := scan(&this.Id, &this.Name, &this.Surname, &this.Birthdate, &this.Email, &this.CreatedAt, &this.UpdatedAt, &this.IsDeleted /*, &this.Username, &this.Password*/)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return error2.NewNotFoundHttpError("user not found")
		}
		return err
	}
	if this.IsDeleted {
		return error2.NewBadRequestHttpError("user has been deleted")
	}
	return nil
}
