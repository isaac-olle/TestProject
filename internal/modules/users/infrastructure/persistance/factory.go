package persistance

import (
	"TestProject/internal/config"
	"TestProject/internal/connections/database"
	"TestProject/internal/modules/users/domain/contracts"
	"TestProject/internal/modules/users/infrastructure/persistance/repositories"
	"encoding/json"
	"errors"
)

// Com que hi ha un repository per llibreria de config, puc inicialitzar cada user repository amb la llibreria que necessiti.
// Es podria afegir una capa d'extraccio més gran per tal de que totes les configs tinguessin alguns metodes en concret i agruparho a traves duna interficie

// De moment aquest 2 repositoris fan el mateix, estan separats perque en un futur potser tenen funcionalitats diferents
func NewUsersRepository(cnf *config.DeviceConfig) (contracts.IUsersRepository, error) {
	switch cnf.Driver {
	case "mysql":
		var mySqlConfig *config.MySqlConfig
		err := json.Unmarshal(cnf.Data, mySqlConfig)
		if err != nil {
			return nil, err
		}
		db, err := database.NewMySqlConnection(mySqlConfig)
		if err != nil {
			return nil, err
		}
		return repositories.NewUserMySqlRepository(db), nil
	case "postgres":
		var postgresConfig *config.PostgresConfig
		err := json.Unmarshal(cnf.Data, &postgresConfig)
		if err != nil {
			return nil, err
		}
		db, err := database.NewPostgresConnection(postgresConfig)
		if err != nil {
			return nil, err
		}
		return repositories.NewUserPostgresRepository(db), nil
	default:
		return nil, errors.New("driver not supported")
	}
}
