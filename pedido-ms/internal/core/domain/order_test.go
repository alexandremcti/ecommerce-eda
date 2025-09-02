package domain_test

import (
	"pedido-ms/internal/core/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	t.Run("Calcula valor total", func(t *testing.T) {
		//Arrange
		orderInput := createOrderInput()
		var expectedTotal float64 = 0.0

		for _, i := range orderInput.Items {
			expectedTotal += (float64(i.Count * int(i.Price)))
		}

		order := domain.CreateOrder(
			*orderInput,
		)
		//Act
		order.TotalAmount()
		//Assert
		assert.Equal(t, expectedTotal, order.TotalAmount())
	})
}

func createOrderInput() *domain.OrderParams {
	ei := domain.EntityIdImp{}

	ei.GenerateId()

	item1 := domain.ItemParams{
		ProductId: ei.ID(),
		Count:     5,
		Price:     5.5,
	}

	ei.GenerateId()
	item2 := domain.ItemParams{
		ProductId: ei.ID(),
		Count:     5,
		Price:     10,
	}
	items := []domain.ItemParams{item1, item2}

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

	orderInput := domain.OrderParams{
		Items:    items,
		Customer: customer,
		Payment:  payment,
	}

	return &orderInput
}
