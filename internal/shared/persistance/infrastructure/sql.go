package infrastructure

import (
	httperror "TestProject/internal/error"
	"TestProject/internal/shared/domain/contracts"
	"TestProject/internal/shared/persistance/domain"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func SqlInsertOrUpdate(db *sql.DB, query string, params ...any) error {
	if !(strings.HasPrefix(query, "UPDATE") || strings.HasPrefix(query, "INSERT")) {
		return httperror.NewInternalServerError("invalid query")
	}
	_, err := db.Exec(query, params...)
	if err != nil {
		return httperror.NewInternalServerError(fmt.Sprintf("error while sending the database query. Error: %s", err.Error()))
	}
	return nil
}

func SqlGetById[M contracts.IDomainEntity, T domain.ISqlModel[M]](db *sql.DB, zero T, query string, params ...any) (M, error) {
	if !strings.HasPrefix(query, "SELECT") {
		return *new(M), errors.New("query must start with 'get'")
	}
	row := db.QueryRow(query, params...)
	if row.Err() != nil {
		return *new(M), httperror.NewInternalServerError(fmt.Sprintf("error sending query into database: %s"))
	}
	err := zero.LoadFromDB(row.Scan)
	if err != nil {
		return *new(M), err
	}
	return zero.ToDomain()
}

func SqlGetAll[M contracts.IDomainEntity, T domain.ISqlModel[M]](db *sql.DB, zero T, query string, params ...any) ([]M, error) {
	var result []M
	rows, err := db.Query(query, params)
	if err != nil {
		return nil, errors.New("error executing insert query: " + err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var domainModel M
		err = zero.LoadFromDB(rows.Scan)
		if err != nil {
			return nil, err
		}
		domainModel, err = zero.ToDomain()

		result = append(result, domainModel)
	}
	return result, nil
}

func SqlDelete(db *sql.DB, query string, params ...any) error {
	if !(strings.HasPrefix(query, "DELETE") || strings.HasPrefix(query, "UPDATE")) {
		return errors.New("query must start with 'DELETE' or 'UPDATE'")
	}
	_, err := db.Exec(query, params...)
	if err != nil {
		return err
	}
	return nil
}
