package entities

import (
	cerror "TestProject/internal/error"
	"TestProject/internal/modules/shared/domain/users"
	"TestProject/internal/modules/users/domain/dtos"
	"TestProject/internal/modules/users/domain/vo"
	"encoding/json"
)

type User struct {
	id        *users.UserId
	name      *vo.UserName
	surname   *vo.UserSurname
	email     *vo.UserEmail
	birthdate *vo.UserBirthdate
	createdAt *vo.UserCreatedAt

	isDeleted bool
}

func NewUserFromUnvaluedObjects(params dtos.UserParams) (*User, error) {
	idVo, err := users.NewUserIdFromString(params.Id)
	cerror.ManageMultipleError(err)
	nameVo, err := vo.NewUserName(params.Name)
	cerror.ManageMultipleError(err)
	surnameVo, err := vo.NewUserSurname(params.Surname)
	cerror.ManageMultipleError(err)
	birthDateVo, err := vo.NewUserBirthdateFromString(params.Birthdate, params.BirthdateTimeFormat)
	cerror.ManageMultipleError(err)
	emailVo, err := vo.NewUserEmail(params.Email)
	cerror.ManageMultipleError(err)
	createdAtVo, err := vo.NewCreatedAtFromString(params.CreatedAt, params.CreatedAtTimeFormat)
	cerror.ManageMultipleError(err)
	if cerror.ErrorSet != nil {
		return nil, cerror.ConculdeMultipleError()
	}
	return &User{
		id:        idVo,
		name:      nameVo,
		surname:   surnameVo,
		email:     emailVo,
		birthdate: birthDateVo,
		createdAt: createdAtVo,
		isDeleted: params.IsDeleted,
	}, nil
}

func (u *User) Id() *users.UserId {
	return u.id
}

func (u *User) Name() *vo.UserName {
	return u.name
}

func (u *User) Surname() *vo.UserSurname {
	return u.surname
}

func (u *User) Email() *vo.UserEmail {
	return u.email
}

func (u *User) Birthdate() *vo.UserBirthdate {
	return u.birthdate
}

func (u *User) CreatedAt() *vo.UserCreatedAt {
	return u.createdAt
}

func (u *User) SetName(name string) (*User, error) {
	nameVo, err := vo.NewUserName(name)
	if err != nil {
		return nil, err
	}
	u.name = nameVo
	return u, nil
}

func (u *User) SetSurname(surname string) (*User, error) {
	surnameVo, err := vo.NewUserSurname(surname)
	if err != nil {
		return nil, err
	}
	u.surname = surnameVo
	return u, nil
}

func (u *User) SetBirthdate(birthdate, format string) (*User, error) {
	birthDateVo, err := vo.NewUserBirthdateFromString(birthdate, format)
	if err != nil {
		return nil, err
	}
	u.birthdate = birthDateVo
	return u, nil
}

func (u *User) SetEmail(email string) (*User, error) {
	emailVo, err := vo.NewUserEmail(email)
	if err != nil {
		return nil, err
	}
	u.email = emailVo
	return u, nil
}

func (u *User) PatchName(name *vo.UserName) *User {
	if name == nil {
		return u
	}
	u.name = name
	return u
}

func (u *User) PatchSurname(surname *vo.UserSurname) *User {
	if surname == nil {
		return u
	}
	u.surname = surname
	return u
}

func (u *User) PatchBirthdate(birthdate *vo.UserBirthdate) *User {
	if birthdate == nil {
		return u
	}
	u.birthdate = birthdate
	return u
}

func (u *User) PatchEmail(email *vo.UserEmail) *User {
	if email == nil {
		return u
	}
	u.email = email
	return u
}

func (u *User) CheckDeleted() error {
	if u.isDeleted {
		return cerror.NewBadRequestHttpError("user is deleted")
	}
	return nil
}

func (u *User) ToResponse(params ...any) any {
	return newJsonUser(u)
}

type jsonUser struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Birthdate string `json:"birthdate"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

func newJsonUser(user *User) *jsonUser {
	return &jsonUser{
		Id:        user.Id().ToString(),
		Name:      user.Name().ToString(),
		Surname:   user.Surname().ToString(),
		Birthdate: user.Birthdate().ToString(),
		Email:     user.Email().ToString(),
		CreatedAt: user.createdAt.ToString(),
	}
}

func (this jsonUser) Serialize() ([]byte, error) {
	return json.Marshal(this)
}
