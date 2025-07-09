package broker

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

var B *RabbitMQBrokerPublisher

type NotifyError struct {
	message string
}

func (e *NotifyError) Error() string {
	return e.message
}

type IBrokerPublisher interface {
	CreatePublishers(bindNames []string) error
	Notify(ctx *context.Context, bindName string, object interface{}) error
}

type RabbitMQBrokerPublisher struct {
	connection *amqp.Connection
	chanel     *amqp.Channel
}

func (b *RabbitMQBrokerPublisher) Notify(ctx *context.Context, bindName string, obj interface{}) error {
	eb, err := json.Marshal(obj)
	if err != nil {
		return &NotifyError{message: "Erro ao transformar objeto"}
	}

	err = b.chanel.PublishWithContext(*ctx,
		bindName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(eb),
		},
	)
	if err != nil {
		return &NotifyError{message: "Erro ao publicar evento " + bindName}

	}
	return nil
}

func CreateConnection(uri string) error {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	B = &RabbitMQBrokerPublisher{
		connection: conn,
		chanel:     ch,
	}

	return nil
}

func (b *RabbitMQBrokerPublisher) CreatePublishers(bindNames []string) error {

	for _, bindName := range bindNames {
		err := b.chanel.ExchangeDeclare(
			bindName,
			"fanout",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *RabbitMQBrokerPublisher) createChannel() (*amqp.Channel, error) {
	ch, err := b.connection.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}
