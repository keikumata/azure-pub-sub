package auth

import (
	amqp "github.com/Azure/azure-amqp-common-go/v3/auth"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/keikumata/azure-pub-sub/internal/aad"
)

type AdalTokenProvider interface {
	GetToken(resource string) (*adal.ServicePrincipalToken, error)
}

type AmqpAuthProviderAdapter struct {
	refresher adal.Refresher
}

func AsJWTTokenProvider(t *adal.ServicePrincipalToken) amqp.TokenProvider {
	// we can safely ignore the error because we provide the token
	provider, _ := aad.NewJWTProvider(aad.JWTProviderWithAADToken(t))
	return provider
}
