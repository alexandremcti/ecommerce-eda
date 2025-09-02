package services

import (
	"context"
	"log"
	"pedido-ms/internal/core/domain"
	"pedido-ms/internal/core/port"
	"pedido-ms/shared/uow"
)

type OrderService struct {
	uow uow.UOW
}

func NewOrderService(uow uow.UOW) *OrderService {
	return &OrderService{
		uow: uow,
	}
}

func (os *OrderService) Create(ctx *context.Context, order *domain.Order) (*domain.Order, error) {
	order.CalculateTotalAmount()

	log.Println("[Order Service] Order:", order)

	err := os.uow.Do(ctx, func(ctx *context.Context, tx uow.TX) error {
		orderRepository, err := uow.GetAs[port.OrderRepository](tx, "OrderRepository")
		if err != nil {
			return err
		}

		err = orderRepository.Create(ctx, order)
		if err != nil {
			return err
		}

		outboxRepository, err := uow.GetAs[port.OutboxRepository](tx, "OutboxRepository")
		if err != nil {
			return err
		}

		ei := domain.EntityIdImp{}
		ei.GenerateId()

		err = outboxRepository.Save(ctx, &domain.OutboxMessage{
			Id:        ei.ID(),
			EventName: "order-created",
			EventData: struct {
				OrderId string
				Order   *domain.OrderParams
			}{
				OrderId: order.Map().ID(),
				Order:   order.Map(),
			},
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println("[Order Service] erro ao criar order: ", err)
		return nil, err
	}

	return order, nil
}
