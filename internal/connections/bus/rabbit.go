package bus

import (
	"TestProject/internal/config"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitConnection(config *config.RabbitMQConfig) (*amqp.Channel, error) {
	if config == nil {
		return nil, errors.New("bus config config is nil")
	}
	conn, err := amqp.Dial("amqp://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/")
	if err != nil {
		return nil, errors.New("connect to RabbitMQ failed: " + err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.New("create channel failed: " + err.Error())
	}
	_, err = ch.QueueDeclare(
		config.QueueName,
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.New("declare queue failed: " + err.Error())
	}
	return ch, nil
}
