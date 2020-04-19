package rest

import (
	"encoding/json"
	"time"

	"github.com/DeKal/bookstore_oath-api/src/domain/users"
	"github.com/DeKal/bookstore_users-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	userRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:9001",
		Timeout: 100 * time.Millisecond,
	}
)

// Repository export interface for Rest call
type Repository interface {
	Login(string, string) (*users.User, *errors.RestError)
}
type repository struct {
}

// NewRepository new rest Repository implementation
func NewRepository() Repository {
	return &repository{}
}

func (*repository) Login(email string, password string) (*users.User, *errors.RestError) {
	request := users.LoginRequest{
		Email:    email,
		Password: password,
	}
	response := userRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid rest client response when trying to get user")
	}
	if response.StatusCode > 299 {
		restError := errors.RestError{}
		if err := json.Unmarshal(response.Bytes(), &restError); err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface return when trying to login user")
		}
		return nil, &restError
	}

	user := users.User{}
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("Error while trying to unmarshal user response")
	}
	return &user, nil
}
