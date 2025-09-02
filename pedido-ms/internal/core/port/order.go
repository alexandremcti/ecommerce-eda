package port

import (
	"context"
	"pedido-ms/internal/core/domain"
)

type OrderRepository interface {
	Create(ctx *context.Context, order *domain.Order) error
}

type OrderService interface {
	Create(ctx *context.Context, order *domain.Order) (*domain.Order, error)
}

type OrderOutput interface {
	OrderCreated(ctx *context.Context, order *domain.Order) error
}

type OutboxRepository interface {
	Save(ctx *context.Context, entity *domain.OutboxMessage) error
}
