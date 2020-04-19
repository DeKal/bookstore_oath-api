package svcaccesstoken

import (
	"strings"

	accesstoken "github.com/DeKal/bookstore_oath-api/src/domain/access_token"
	"github.com/DeKal/bookstore_oath-api/src/repository/db"
	"github.com/DeKal/bookstore_oath-api/src/repository/rest"
	"github.com/DeKal/bookstore_users-api/src/utils/errors"
)

// Service service interface
type Service interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestError)
	Create(*accesstoken.Request) (*accesstoken.AccessToken, *errors.RestError)
	UpdateExpirationTime(*accesstoken.AccessToken) *errors.RestError
}

type service struct {
	dbRepository   db.Repository
	userRepository rest.Repository
}

// NewService return service implementation
func NewService(dbRepo db.Repository, userRepo rest.Repository) Service {
	return &service{
		dbRepository:   dbRepo,
		userRepository: userRepo,
	}
}

// GetByID return token by user id
func (s *service) GetByID(accessTokenID string) (*accesstoken.AccessToken, *errors.RestError) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token id")
	}
	accessToken, err := s.dbRepository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request *accesstoken.Request) (*accesstoken.AccessToken, *errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := s.userRepository.Login(request.UserName, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := accesstoken.GetNewAccessToken(user.ID)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepository.Create(at); err != nil {
		return nil, err
	}
	return at, nil
}

func (s *service) UpdateExpirationTime(token *accesstoken.AccessToken) *errors.RestError {
	if err := token.Validate(); err != nil {
		return err
	}

	return s.dbRepository.UpdateExpirationTime(token)
}
