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

	OutboxMockRepositoryArgs struct {
		ctx    *context.Context
		entity *domain.OutboxMessage
	}
	OrderMockRepository struct {
		args   OrderMockRepositoryArgs
		orders []*domain.Order
	}

	OutboxMockRepository struct {
		args     OutboxMockRepositoryArgs
		entities []*domain.OutboxMessage
	}
)

func NewOrderRepository() *OrderMockRepository {
	r := &OrderMockRepository{}
	return r
}

func NewOutboxRepository() *OutboxMockRepository {
	r := &OutboxMockRepository{}
	return r
}

func (r *OrderMockRepository) Create(ctx *context.Context, o *domain.Order) error {
	r.args.ctx = ctx
	r.args.o = o
	r.orders = append(r.orders, o)

	return nil
}

func (r *OutboxMockRepository) Save(ctx *context.Context, entity *domain.OutboxMessage) error {
	r.args.ctx = ctx
	r.args.entity = entity
	r.entities = append(r.entities, entity)

	return nil
}
