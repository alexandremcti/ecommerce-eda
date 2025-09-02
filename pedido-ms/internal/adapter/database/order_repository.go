package database

import (
	"context"
	"fmt"
	"pedido-ms/internal/core/domain"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	OrderRepository struct {
		db   *mongo.Client
		coll *mongo.Collection
	}

	ItemEntinty struct {
		ProductId string  `bson:"productId"`
		Count     int     `bson:"count"`
		Price     float64 `bson:"price"`
	}

	AddressEntity struct {
		Street     string `bson:"street"`
		Number     string `bson:"number"`
		PostalCode string `bson:"postalCode"`
		City       string `bson:"city"`
		State      string `bson:"state"`
	}

	CustomerEntity struct {
		Id              string        `bson:"id"`
		FirstName       string        `bson:"firstName"`
		LastName        string        `bson:"lastName"`
		Email           string        `bson:"email"`
		DeliveryAddress AddressEntity `bson:"deliveryAddress"`
	}

	PaymentEntity struct {
		CardId          string `bson:"cardId"`
		Bin             string `bson:"bin"`
		NumToken        string `bson:"numToken"`
		CardholderName  string `bson:"cardhoderName"`
		SecurityCode    string `bson:"securityCode"`
		ExpirationMonth string `bson:"expirationMonth"`
		ExpirationYear  string `bson:"expirationYear"`
		Brand           string `bson:"brand"`
	}

	OrderEntity struct {
		Id          string         `bson:"orderId"`
		TotalAmount float64        `bson:"totalAmount"`
		OrderStatus string         `bson:"orderStatus"`
		Items       []ItemEntinty  `bson:"items"`
		Customer    CustomerEntity `bson:"customer"`
		Payment     PaymentEntity  `bson:"payment"`
		Qualified   bool           `bson:"qualified"`
		Reserved    bool           `bson:"reserved"`
		CreatedAt   time.Time      `bson:"createdAt"`
		UpdatedAt   time.Time      `bson:"updatedAt"`
	}
)

func CreateOrderRepository() *OrderRepository {
	r := OrderRepository{
		db:   Client,
		coll: Client.Database("pedidos_db").Collection("orders"),
	}
	return &r
}

func (r *OrderRepository) Create(ctx *context.Context, o *domain.Order) error {

	e := r.mapToORM(o)

	fmt.Println("[Order Repository] Criando order: ", e)

	_, err := r.coll.InsertOne(*ctx, e)

	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) mapToORM(o *domain.Order) *OrderEntity {
	e := OrderEntity{}
	mappiedOrder := o.Map()
	e.Id = o.ID()
	e.CreatedAt = mappiedOrder.CreatedAt
	e.UpdatedAt = mappiedOrder.UpdatedAt
	e.OrderStatus = mappiedOrder.OrderStatus.String()
	e.Qualified = mappiedOrder.Qualified
	e.Reserved = mappiedOrder.Reserved
	e.TotalAmount = mappiedOrder.TotalAmount

	items := []ItemEntinty{}
	for _, i := range mappiedOrder.Items {
		items = append(items, ItemEntinty{ProductId: i.ProductId, Count: i.Count, Price: i.Price})
	}
	e.Items = items

	e.Customer.Id = mappiedOrder.Customer.Id
	e.Customer.FirstName = mappiedOrder.Customer.FirstName
	e.Customer.LastName = mappiedOrder.Customer.LastName
	e.Customer.Email = mappiedOrder.Customer.Email
	e.Customer.DeliveryAddress.City = mappiedOrder.Customer.City
	e.Customer.DeliveryAddress.Number = mappiedOrder.Customer.Streetnumber
	e.Customer.DeliveryAddress.PostalCode = mappiedOrder.Customer.PostalCode
	e.Customer.DeliveryAddress.State = mappiedOrder.Customer.State
	e.Customer.DeliveryAddress.Street = mappiedOrder.Customer.Street

	e.Payment.Bin = mappiedOrder.Payment.Bin
	e.Payment.CardId = mappiedOrder.Payment.CardId
	e.Payment.Brand = mappiedOrder.Payment.Brand.String()
	e.Payment.CardholderName = mappiedOrder.Payment.CardholderName
	e.Payment.ExpirationMonth = mappiedOrder.Payment.ExpirationMonth
	e.Payment.ExpirationYear = mappiedOrder.Payment.ExpirationYear
	e.Payment.NumToken = mappiedOrder.Payment.NumToken
	e.Payment.SecurityCode = mappiedOrder.Payment.SecurityCode

	return &e
}
