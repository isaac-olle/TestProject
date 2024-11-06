package bus

import (
	"TestProject/internal/config"
	"errors"
	"github.com/redis/go-redis/v9"
)

type RedisBusClient struct {
	*redis.Client
	channel string
}

func NewRedisConection(config *config.RedisConfig) (*RedisBusClient, error) {
	if config == nil {
		return nil, errors.New("redis config is nil")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Username: config.User,
		Password: config.Password,
	})
	return &RedisBusClient{
		Client:  rdb,
		channel: config.QueueName,
	}, nil
}

func (this *RedisBusClient) Channel() string {
	return this.channel
}
