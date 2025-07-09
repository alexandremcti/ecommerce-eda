package services

import (
	"context"
	"log"
	"pedido-ms/internal/core/domain"
	"pedido-ms/internal/core/port"
)

type OrderService struct {
	orderRepository port.OrderRepository
	orderOuput      port.OrderOutput
}

func NewOrderService(orderRepository port.OrderRepository, orderOuput port.OrderOutput) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		orderOuput:      orderOuput,
	}
}

func (os *OrderService) Create(ctx *context.Context, order *domain.Order) (*domain.Order, error) {
	order.CalculateTotalAmount()

	log.Println("[Order Service] Order:", order)

	err := os.orderRepository.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	err = os.orderOuput.OrderCreated(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
