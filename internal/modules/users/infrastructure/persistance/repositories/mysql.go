package repositories

import (
	"TestProject/internal/modules/shared/domain/users"
	"TestProject/internal/modules/users/domain/entities"
	entities2 "TestProject/internal/modules/users/infrastructure/persistance/entities"
	mysql "TestProject/internal/shared/persistance/infrastructure"
	"database/sql"
)

type UserMySqlRepository struct {
	db *sql.DB
}

func NewUserMySqlRepository(db *sql.DB) *UserMySqlRepository {
	return &UserMySqlRepository{db: db}
}

func (this *UserMySqlRepository) Create(user *entities.User) error {
	const query = "INSERT INTO users (name, surname, birthdate, email) VALUES (?, ?, ?, ?) RETURNING *"
	return mysql.SqlInsertOrUpdate(this.db, query, user.Name().ToString(), user.Surname().ToString(), user.Birthdate().Value(), user.Email().ToString())
}

func (this *UserMySqlRepository) GetById(id *users.UserId) (*entities.User, error) {
	const query = "SELECT * FROM users WHERE id = ?"
	return mysql.SqlGetById[*entities.User, *entities2.UserSQLDatabaseTable](this.db, entities2.NewEmptyUserMySQLDatabaseTable(), query, id)
}

func (this *UserMySqlRepository) GetAll() ([]*entities.User, error) {
	const query = "SELECT * FROM users"
	return mysql.SqlGetAll[*entities.User, *entities2.UserSQLDatabaseTable](this.db, entities2.NewEmptyUserMySQLDatabaseTable(), query)
}

func (this *UserMySqlRepository) Update(user *entities.User, id *users.UserId) error {
	const query = "UPDATE users SET name = ?, surname = ?, birthdate = ?, email = ? WHERE id = ? RETURNING *"
	return mysql.SqlInsertOrUpdate(this.db, query, user.Name().ToString(), user.Surname().ToString(), user.Birthdate().Value(), user.Email().ToString(), id)
}

func (this *UserMySqlRepository) Delete(id string) error {
	const query = "UPDATE users SET isDeleted = true WHERE id = ?"
	return mysql.SqlDelete(this.db, query, id)
}
