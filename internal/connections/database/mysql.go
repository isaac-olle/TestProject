package database

import (
	"TestProject/internal/config"
	"database/sql"
	"errors"
)

var mySqlConnection *sql.DB

func NewMySqlConnection(config *config.MySqlConfig) (*sql.DB, error) {
	if mySqlConnection != nil {
		return mySqlConnection, nil
	}
	if config == nil {
		return nil, errors.New("mysql config is nil")
	}
	db, err := sql.Open("mysql", config.User+":"+config.Password+"@tcp("+config.Host+":"+config.Port+")/"+config.Database)
	if err != nil {
		return nil, errors.New("failed to connect to mysql: " + err.Error())
	}

	mySqlConnection = db

	return mySqlConnection, nil
}
