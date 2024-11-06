package error

import (
	"TestProject/internal/config"
	errorConnections "TestProject/internal/connections/error"
	errorDomain "TestProject/internal/shared/bus/domain/error"
	"TestProject/internal/shared/bus/infrastructure/error/repositories"
	"encoding/json"
	"errors"
)

func NewErrorRepository(cnf *config.DeviceConfig) (errorDomain.IErrorRepository, error) {
	switch cnf.Driver {
	case "dynamoDb":
		var dynamoDbConfig *config.DynamoDbConfig
		err := json.Unmarshal(cnf.Data, &dynamoDbConfig)
		if err != nil {
			return nil, errors.New("invalid config for dynamoDb driver: " + err.Error())
		}
		client, err := errorConnections.NewDynamoDBConnection(dynamoDbConfig)
		if err != nil {
			return nil, err
		}
		return repositories.NewDynamoDbErrorRepository(client), nil
	default:
		return nil, errors.New("invalid driver for repository")
	}
}
