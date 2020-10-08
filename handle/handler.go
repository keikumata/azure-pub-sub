package handle

import (
	"context"

	servicebus "github.com/Azure/azure-service-bus-go"
)

// Handle is a func to handle the message received from a subscription
type Handle func(ctx context.Context, message, messageType string) Handler

func (h Handle) Do(ctx context.Context, _ Handler, message *servicebus.Message) Handler {
	val, ok := message.UserProperties["type"]
	// we require the message to have a type set on user properties
	if !ok {
		return Abandon()
	}
	return h(ctx, string(message.Data), val.(string))
}

type Handler interface {
	Do(ctx context.Context, orig Handler, message *servicebus.Message) Handler
}
