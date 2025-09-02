package database

import (
	"context"
	"fmt"
	"pedido-ms/internal/core/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	OutboxRepository struct {
		db   *mongo.Client
		coll *mongo.Collection
	}
)

func NewOutboxRepository() *OutboxRepository {
	return &OutboxRepository{
		db:   Client,
		coll: Client.Database("pedidos_db").Collection("outbox"),
	}
}

func (o *OutboxRepository) Save(ctx *context.Context, entity *domain.OutboxMessage) error {
	fmt.Println("[Outbox Repository] Criando order", *entity)

	bytes, _ := bson.Marshal(entity.EventData)
	_ = bson.Unmarshal(bytes, &entity.Payload)

	fmt.Println("[Outbox Repository] Order formatada para bson:", entity.Payload)

	_, err := o.coll.InsertOne(*ctx, bson.D{
		bson.E{Key: "correlation_id", Value: entity.Id},
		bson.E{Key: "event_name", Value: entity.EventName},
		bson.E{Key: "payload", Value: entity.Payload},
	})
	if err != nil {
		return err
	}

	return nil
}
