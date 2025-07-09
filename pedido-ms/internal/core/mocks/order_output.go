package mocks

import (
	"context"
	"pedido-ms/internal/core/domain"
)

type OrderOutputMock struct {
	Input *domain.Order
}

func NewOrderOutput() *OrderOutputMock {
	return &OrderOutputMock{}
}

func (oom *OrderOutputMock) OrderCreated(ctx *context.Context, order *domain.Order) error {
	oom.Input = order
	return nil
}
