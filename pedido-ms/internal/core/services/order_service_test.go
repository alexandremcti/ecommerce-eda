package services_test

import (
	"context"
	"pedido-ms/internal/core/mocks"
	"pedido-ms/internal/core/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	ctx := context.Background()

	t.Run("Cria pedido", func(t *testing.T) {
		order := mocks.CreateOrderEntity()
		orderRepository := mocks.NewOrderRepository()
		outboxRepository := mocks.NewOutboxRepository()
		uow := mocks.NewUnitOfWorkMock()
		uow.Register("OrderRepository", orderRepository)
		uow.Register("OutboxRepository", outboxRepository)

		service := services.NewOrderService(uow)
		orderResult, err := service.Create(&ctx, order)

		assert.Nil(t, err)
		assert.Contains(t, orderResult.ID(), order.ID())
		assert.ObjectsAreEqualValues(orderResult.Map(), order.Map())
	})
}
