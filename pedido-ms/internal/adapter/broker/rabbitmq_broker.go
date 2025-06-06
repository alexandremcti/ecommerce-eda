package broker

import (
	"pedido-ms/internal/core/domain"

	"github.com/wagslane/go-rabbitmq"
)

var (
	conn *rabbitmq.Conn
	err  error
)

type IBrokerPublisher interface {
	Notify(bindName string, order domain.Order)
}

type BrokerPublisher struct {
	publisher *rabbitmq.Publisher
}

func (b *BrokerPublisher) Notify(order domain.Order) {
	//b.publisher.Publish([]byte(order))
}

func CreatePublisher(bindName string) (*BrokerPublisher, error) {
	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(bindName),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)

	if err != nil {
		return nil, err
	}

	return &BrokerPublisher{publisher: publisher}, nil
}

func CreateConnection(uri string) (*rabbitmq.Conn, error) {
	conn, err = rabbitmq.NewConn(
		uri,
		rabbitmq.WithConnectionOptionsLogging,
	)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
