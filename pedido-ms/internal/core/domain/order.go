package domain

import (
	"pedido-ms/internal/core/enum"
	"time"
)

type (
	OrderId struct {
		*EntityIdImp
	}

	OrderParams struct {
		OrderId
		TotalAmount float64
		OrderStatus enum.OrderState
		Items       []ItemParams
		Customer    CustomerParams
		Payment     PaymentParams
		Qualified   bool
		Reserved    bool
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	Order struct {
		OrderId
		totalAmount float64
		orderStatus enum.OrderState
		items       []Item
		customer    *Customer
		payment     *Payment
		qualified   bool
		reserved    bool
		createdAt   time.Time
		updatedAt   time.Time
	}
)

func CreateOrder(input OrderParams) *Order {
	ei := EntityIdImp{}
	ei.GenerateId()
	input.OrderId = OrderId{EntityIdImp: &ei}

	order := Order{
		OrderId:     input.OrderId,
		orderStatus: enum.Recebido,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
		qualified:   false,
		reserved:    false,
		customer:    CreateCustomer(input.Customer),
		items:       bindItems(input.Items),
		payment:     CreatePayment(input.Payment),
	}

	return &order
}

func RecoverOrder(input OrderParams) *Order {

	order := Order{
		OrderId:     input.OrderId,
		totalAmount: input.TotalAmount,
		orderStatus: input.OrderStatus,
		createdAt:   input.CreatedAt,
		updatedAt:   input.UpdatedAt,
		qualified:   input.Qualified,
		reserved:    input.Reserved,
		customer:    RecoverCustomer(input.Customer),
		items:       bindItems(input.Items),
		payment:     RecoverPayment(input.Payment),
	}

	return &order
}

func bindItems(input []ItemParams) []Item {
	items := []Item{}
	for _, i := range input {
		item := CreateItem(i)
		items = append(items, item)
	}
	return items
}

func (o *Order) CalculateTotalAmount() {
	ta := 0.0
	for _, i := range o.items {
		ta += float64(i.count * int(i.price))
	}
	o.totalAmount = ta
}

func (o *Order) TotalAmount() float64 {
	return o.totalAmount
}

func (o *Order) ID() string {
	return o.OrderId.id
}

func (o *Order) Map() *OrderParams {
	op := OrderParams{}
	op.OrderId = o.OrderId
	op.TotalAmount = o.totalAmount
	op.OrderStatus = o.orderStatus
	op.Reserved = o.reserved
	op.Qualified = o.qualified
	op.Customer = *o.customer.Map()
	op.Payment = *o.payment.Map()
	op.CreatedAt = o.createdAt
	op.UpdatedAt = o.updatedAt
	op.Items = []ItemParams{}
	for _, i := range o.items {
		op.Items = append(op.Items, *i.Map())
	}

	return &op
}
