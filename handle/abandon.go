package handle

import (
	"context"

	servicebus "github.com/Azure/azure-service-bus-go"
)

func Abandon() Handler {
	return &abandon{}
}

type abandon struct {
}

func (a *abandon) Do(ctx context.Context, _ Handler, message *servicebus.Message) Handler {
	message.Abandon(ctx)
	return done()
}
