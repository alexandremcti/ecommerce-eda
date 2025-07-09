package dto

import (
	"log"
	"pedido-ms/internal/core/domain"
	"pedido-ms/internal/core/enum"
	"time"
)

type (
	ItemDTO struct {
		ProductId string  `json:"productId" binding:"required"`
		Count     int     `json:"count" binding:"required"`
		Price     float64 `json:"price" binding:"required"`
	}

	DeliveryAddressDTO struct {
		Street     string `json:"street" binding:"required"`
		Number     string `json:"number" binding:"required"`
		PostalCode string `json:"postalCode" binding:"required"`
		City       string `json:"city" binding:"required"`
		State      string `json:"state" binding:"required"`
	}

	CustomerDTO struct {
		Id              string             `json:"id" binding:"required"`
		FirstName       string             `json:"first_name" binding:"required"`
		LastName        string             `json:"last_name" binding:"required"`
		Email           string             `json:"email" binding:"required"`
		DeliveryAddress DeliveryAddressDTO `json:"deliveryAddress" binding:"required"`
	}

	PaymentDTO struct {
		CardId          string `json:"cardId" binding:"required"`
		Bin             string `json:"bin" binding:"required"`
		NumToken        string `json:"number_token" binding:"required"`
		CardholderName  string `json:"cardholder_name" binding:"required"`
		SecurityCode    string `json:"security_code" binding:"required"`
		ExpirationMonth string `json:"expiration_month" binding:"required"`
		ExpirationYear  string `json:"expiration_year" binding:"required"`
		Brand           string `json:"brand" binding:"required"`
	}

	OrderDTO struct {
		Id          string      `json:"id,omitempty"`
		TotalAmount float64     `json:"totalAmount,omitempty"`
		OrderStatus string      `json:"orderStatus,omitempty"`
		Qualified   bool        `json:"qualified,omitempty"`
		Reserved    bool        `json:"reserved,omitempty"`
		CreatedAt   time.Time   `json:"createdAt,omitempty"`
		UpdatedAt   time.Time   `json:"updatedAt,omitempty"`
		Items       []ItemDTO   `json:"items" binding:"required"`
		Customer    CustomerDTO `json:"customer" binding:"required"`
		Payment     PaymentDTO  `json:"payment" binding:"required"`
	}
)

func (odto *OrderDTO) ToEntity() *domain.Order {
	input := domain.OrderParams{
		Qualified: false,
		Reserved:  false,
		Customer:  odto.Customer.toEntityInput(),
		Items:     bindItems(odto.Items),
		Payment:   odto.Payment.toEntityInput(),
	}

	log.Println("[ORDER DTO] order input params:", input)

	order := domain.CreateOrder(input)

	log.Println("[ORDER DTO] order entity:", order)

	return order
}

func (cdto *CustomerDTO) ToEntity() *domain.Customer {
	input := cdto.toEntityInput()
	customer := domain.CreateCustomer(input)
	return customer
}

func (cdto *CustomerDTO) toEntityInput() domain.CustomerParams {
	input := domain.CustomerParams{
		FirstName:    cdto.FirstName,
		LastName:     cdto.LastName,
		Email:        cdto.Email,
		Street:       cdto.DeliveryAddress.Street,
		Streetnumber: cdto.DeliveryAddress.Number,
		PostalCode:   cdto.DeliveryAddress.PostalCode,
		City:         cdto.DeliveryAddress.City,
		State:        cdto.DeliveryAddress.State,
	}
	return input
}

func bindItems(input []ItemDTO) []domain.ItemParams {
	items := []domain.ItemParams{}
	for _, i := range input {
		item := domain.ItemParams{
			ProductId: i.ProductId,
			Count:     i.Count,
			Price:     i.Price,
		}
		items = append(items, item)
	}
	return items
}

func (pdto *PaymentDTO) ToEntity() *domain.Payment {
	input := pdto.toEntityInput()
	payment := domain.CreatePayment(input)

	return payment
}

func (pdto *PaymentDTO) toEntityInput() domain.PaymentParams {
	input := domain.PaymentParams{
		Brand:           enum.GetBrandBy(pdto.Brand),
		CardId:          pdto.CardId,
		Bin:             pdto.Bin,
		NumToken:        pdto.NumToken,
		CardholderName:  pdto.CardholderName,
		SecurityCode:    pdto.SecurityCode,
		ExpirationMonth: pdto.ExpirationMonth,
		ExpirationYear:  pdto.ExpirationYear,
	}

	return input
}

func ToDTO(o *domain.OrderParams) *OrderDTO {
	order := OrderDTO{
		Id:          o.ID(),
		TotalAmount: o.TotalAmount,
		OrderStatus: o.OrderStatus.String(),
		Qualified:   o.Qualified,
		Reserved:    o.Reserved,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
	}

	order.Customer = CustomerDTO{
		Id:        o.Customer.Id,
		FirstName: o.Customer.FirstName,
		LastName:  o.Customer.LastName,
		Email:     o.Customer.Email,
		DeliveryAddress: DeliveryAddressDTO{
			Street:     o.Customer.Street,
			Number:     o.Customer.Streetnumber,
			PostalCode: o.Customer.PostalCode,
			City:       o.Customer.Street,
			State:      o.Customer.State,
		},
	}

	order.Items = []ItemDTO{}
	for _, i := range o.Items {
		item := ItemDTO{
			ProductId: i.ProductId,
			Count:     i.Count,
			Price:     i.Price,
		}
		order.Items = append(order.Items, item)
	}

	order.Payment = PaymentDTO{
		CardId:          o.Payment.CardId,
		Bin:             o.Payment.Bin,
		NumToken:        o.Payment.NumToken,
		CardholderName:  o.Payment.CardholderName,
		SecurityCode:    o.Payment.SecurityCode,
		ExpirationMonth: o.Payment.ExpirationMonth,
		ExpirationYear:  o.Payment.ExpirationYear,
		Brand:           o.Payment.Brand.String(),
	}

	return &order
}
