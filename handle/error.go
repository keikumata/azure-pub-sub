package handle

import (
	"context"

	servicebus "github.com/Azure/azure-service-bus-go"
)

func Error(e error) Handler {
	return &errorHandler{err: e}
}

type errorHandler struct {
	err error
}

func (a *errorHandler) Do(ctx context.Context, orig Handler, message *servicebus.Message) Handler {
	// can trace the error
	return done()
}
