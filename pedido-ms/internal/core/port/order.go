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
