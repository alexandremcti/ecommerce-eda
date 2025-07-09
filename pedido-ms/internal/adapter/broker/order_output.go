package broker

import (
	"context"
	"pedido-ms/internal/adapter/dto"
	"pedido-ms/internal/core/domain"
)

type (
	OrderOuputImp struct {
		broker IBrokerPublisher
	}
)

func CreateOutputImp(broker IBrokerPublisher) *OrderOuputImp {
	return &OrderOuputImp{broker: broker}
}

func (ooi *OrderOuputImp) OrderCreated(ctx *context.Context, order *domain.Order) error {
	err := ooi.broker.Notify(ctx, "pedido-criado-out-0", dto.ToDTO(order.Map()))
	if err != nil {
		return err
	}
	return nil
}
