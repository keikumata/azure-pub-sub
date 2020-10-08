package handle

import (
	"context"

	servicebus "github.com/Azure/azure-service-bus-go"
)

func Complete() Handler {
	return &complete{}
}

type complete struct {
}

func (a *complete) Do(ctx context.Context, _ Handler, message *servicebus.Message) Handler {
	if err := message.Complete(ctx); err != nil {
		return Error(err)
	}
	return done()
}
