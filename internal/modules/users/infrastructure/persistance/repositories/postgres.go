package repositories

import (
	"TestProject/internal/modules/users/domain/entities"
	entities2 "TestProject/internal/modules/users/infrastructure/persistance/entities"
	postgres "TestProject/internal/shared/persistance/infrastructure"
	"database/sql"
)

type UserPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (this *UserPostgresRepository) Create(user *entities.User) error {
	const query = "INSERT INTO users (id,name, surname, birthdate, email) VALUES ($1, $2 ,$3, $4, $5)"
	return postgres.SqlInsertOrUpdate(this.db, query, user.Id().ToString(), user.Name().ToString(), user.Surname().ToString(), user.Birthdate().Value(), user.Email().ToString())
}

func (this *UserPostgresRepository) GetById(id string) (*entities.User, error) {
	const query = "SELECT * FROM users WHERE id = $1"
	return postgres.SqlGetById[*entities.User, *entities2.UserSQLDatabaseTable](this.db, entities2.NewEmptyUserMySQLDatabaseTable(), query, id)
}

func (this *UserPostgresRepository) GetAll() ([]*entities.User, error) {
	const query = "SELECT * FROM users"
	return postgres.SqlGetAll[*entities.User, *entities2.UserSQLDatabaseTable](this.db, entities2.NewEmptyUserMySQLDatabaseTable(), query)
}

func (this *UserPostgresRepository) Update(user *entities.User, id string) error {
	const query = "UPDATE users SET name = $1, surname = $2, birthdate = $3, email = $4 WHERE id = $5 RETURNING *"
	return postgres.SqlInsertOrUpdate(this.db, query, user.Name().ToString(), user.Surname().ToString(), user.Birthdate().Value(), user.Email().ToString(), id)
}

func (this *UserPostgresRepository) Delete(id string) error {
	const query = "UPDATE users SET isDeleted = true WHERE id = $1"
	return postgres.SqlDelete(this.db, query, id)
}
