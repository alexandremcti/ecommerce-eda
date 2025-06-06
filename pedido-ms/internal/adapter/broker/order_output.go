package broker

import (
	"time"
)

type (
	ItemEvent struct {
		ProductId string  `json:"productId"`
		Count     int     `json:"count"`
		Price     float64 `json:"price"`
	}

	CustomerEvent struct {
		Id              uint64       `json:"id"`
		FirstName       string       `json:"firstName"`
		LastName        string       `json:"lastName"`
		Email           string       `json:"email"`
		DeliveryAddress AddressEvent `json:"deliveryAdress"`
	}

	AddressEvent struct {
		Street     string `json:"street"`
		Number     string `json:"number"`
		PostalCode string `json:"postaCode"`
		City       string `json:"city"`
		State      string `json:"state"`
	}

	OrderEvent struct {
		Id          uint64      `json:"id"`
		TotalAmount float64     `json:"totalAmount"`
		OrderStatus string      `json:"orderStatus"`
		Items       []ItemEvent `json:"items"`
		Customer    string      `json:"customer"`
		Payment     interface{} `json:"payment"`
		Qualified   bool        `json:"qualified"`
		Reserved    bool        `json:"reserved"`
		CreatedAt   time.Time   `json:"createdAt"`
		UpdatedAt   time.Time   `json:"updatedAt"`
	}
)
