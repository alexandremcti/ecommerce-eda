package mapper

import (
	"pedido-ms/internal/adapter/broker"
	"pedido-ms/internal/adapter/dto"
	"pedido-ms/internal/core/domain"
)

type (
	OrderMapper interface {
		toDTO(order domain.Order) dto.OrderDTO
		toEntity(orderDto dto.OrderDTO) domain.Order
		toEvent(order domain.Order) broker.OrderEvent
	}
)
