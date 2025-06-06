package mocks

import (
	"context"
	"pedido-ms/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

type NewRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

func (r *Repository) Create(ctx *context.Context, o *domain.Order) error {
	args := r.Called(ctx, o)

	return args.Error(0)

}

func NewOrderRepository(t NewRepositoryT) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
