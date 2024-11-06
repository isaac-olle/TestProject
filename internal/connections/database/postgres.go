package database

import (
	"TestProject/internal/config"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

var postgresConnection *sql.DB

func NewPostgresConnection(config *config.PostgresConfig) (*sql.DB, error) {
	if postgresConnection != nil {
		return postgresConnection, nil
	}
	if config == nil {
		return nil, errors.New("postgres config is nil")
	}
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database))
	if err != nil {
		return nil, errors.New("failed to connect to postgres: " + err.Error())
	}
	postgresConnection = db

	return postgresConnection, nil
}
