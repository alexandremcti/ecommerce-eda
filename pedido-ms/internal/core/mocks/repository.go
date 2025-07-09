package mocks

import (
	"context"
	"pedido-ms/internal/core/domain"
)

type (
	OrderMockRepositoryArgs struct {
		ctx *context.Context
		o   *domain.Order
	}
	OrderMockRepository struct {
		args   OrderMockRepositoryArgs
		orders []*domain.Order
	}
)

func NewOrderRepository() *OrderMockRepository {
	r := &OrderMockRepository{}
	return r
}

func (r *OrderMockRepository) Create(ctx *context.Context, o *domain.Order) error {
	r.args.ctx = ctx
	r.args.o = o
	r.orders = append(r.orders, o)

	return nil
}
