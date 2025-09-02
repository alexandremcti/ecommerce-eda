package domain

type OutboxMessage struct {
	Id        string         `bson:"correlation_id"`
	EventName string         `bson:"event_name"`
	Payload   map[string]any `bson:"inline"`
	EventData any
}
