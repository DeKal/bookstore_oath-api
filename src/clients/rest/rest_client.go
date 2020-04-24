package rest

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/mercadolibre/golang-restclient/rest"
)

func init() {
	godotenv.Load()
}

// NewRestClient return new rest client to connect to users api
func NewRestClient() *rest.RequestBuilder {
	return &rest.RequestBuilder{
		BaseURL: getUsersRestHost(),
		Timeout: 100 * time.Millisecond,
	}
}
