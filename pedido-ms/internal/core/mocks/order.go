package mocks

import "pedido-ms/internal/core/domain"

func CreateOrderEntity() *domain.Order {
	item1 := domain.ItemParams{
		ProductId: "any_id",
		Count:     5,
		Price:     5.5,
	}
	item2 := domain.ItemParams{
		ProductId: "any_id2",
		Count:     5,
		Price:     10,
	}
	items := []domain.ItemParams{item1, item2}

	var expectedTotal float64 = 0.0

	for _, i := range items {
		expectedTotal += (float64(i.Count * int(i.Price)))
	}

	ei := domain.EntityIdImp{}

	ei.GenerateId()
	customer := domain.CustomerParams{
		Id:           ei.ID(),
		FirstName:    "Alexandre",
		LastName:     "Carvalho",
		Email:        "alexandre@fakemail.com",
		Street:       "any street",
		Streetnumber: "123",
		PostalCode:   "123456789",
		City:         "Sao Paulo",
		State:        "SP",
	}

	payment := domain.PaymentParams{
		CardId:          "any_cardId",
		Bin:             "any_bind",
		NumToken:        "any_numToke",
		CardholderName:  "Any Name",
		SecurityCode:    "any_code",
		ExpirationMonth: "01",
		ExpirationYear:  "2050",
	}

	order := domain.CreateOrder(
		domain.OrderParams{
			Items:    items,
			Customer: customer,
			Payment:  payment,
		},
	)

	return order
}
