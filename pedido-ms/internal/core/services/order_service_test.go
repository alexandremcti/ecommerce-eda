package services_test

import (
	"context"
	"pedido-ms/internal/core/mocks"
	"pedido-ms/internal/core/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateOrder(t *testing.T) {
	ctx := context.Background()
	t.Run("Cria pedido", func(t *testing.T) {
		order := mocks.CreateOrderEntity()
		orderRepository := mocks.NewOrderRepository(t)

		orderRepository.On("Create").Return(nil)

		service := services.NewOrderService(orderRepository)
		orderResult, err := service.Create(&ctx, order)

		assert.Nil(t, err)
		assert.Contains(t, orderResult, orderResult.ID())
	})
}
