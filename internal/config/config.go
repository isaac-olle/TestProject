package config

import (
	"encoding/json"
	"errors"
	"os"
)

const (
	DEFAULT_CONFIG_PATH = "./config/config_local.json"
)

type MainConfig struct {
	ServerConfig        *BasicConfig  `json:"server"`
	CommandBusConfig    *DeviceConfig `json:"commandBus"`
	QueryBusConfig      *DeviceConfig `json:"queryBus"`
	DatabaseConfig      *DeviceConfig `json:"database"`
	ErrorDatabaseConfig *DeviceConfig `json:"errorDatabase"`
}

type DeviceConfig struct {
	Driver string          `json:"driver"`
	Data   json.RawMessage `json:"data,omitempty"`
}

type BasicConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type DynamoDbConfig struct {
	*BasicConfig
	Protocol  string `json:"protocol"`
	Region    string `json:"region"`
	AccessKey string `json:"accessKey"`
	KeyId     string `json:"keyId"`
}

type RabbitMQConfig struct {
	*BasicConfig
	QueueName string `json:"queueName"`
	User      string `json:"user"`
	Password  string `json:"password"`
}

type RedisConfig struct {
	*BasicConfig
	Password  string `json:"password"`
	User      string `json:"user"`
	QueueName string `json:"queueName"`
}

type MongoConfig struct {
	*BasicConfig
	Database string                 `json:"database"`
	Options  map[string]interface{} `json:"options,omitempty"`
}

type PostgresConfig struct {
	*BasicConfig
	Database string `json:"database"`
	User     string `json:"userName"`
	Password string `json:"password,omitempty"`
}

type MySqlConfig struct {
	*BasicConfig
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password,omitempty"`
}

func SetConfigPath() string {
	if os.Getenv("CONFIG_PATH") != "" {
		return os.Getenv("CONFIG_PATH")
	}
	return DEFAULT_CONFIG_PATH
}

// Aqui hi ha una llibreria de go que fa coses semblants, pero no estalviaria res de codi.
func GetConfig(configPath string) (*MainConfig, error) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, errors.New("Error reading config file: " + err.Error())
	}
	var config *MainConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, errors.New("Error parsing config file: " + err.Error())
	}
	return config, nil
}
