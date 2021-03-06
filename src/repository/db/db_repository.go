package db

import (
	accesstoken "github.com/DeKal/bookstore_oath-api/src/domain/access_token"
	"github.com/DeKal/bookstore_utils-go/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken       = "SELECT access_token, user_id, client_id, expired FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken    = "INSERT INTO access_tokens(access_token, user_id, client_id, expired) VALUES (?,?,?,?);"
	queryUpdateExpirationTime = "UPDATE access_tokens SET expired=? WHERE access_token=?;"
)

// Repository service interface
type Repository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestError)
	Create(*accesstoken.AccessToken) *errors.RestError
	UpdateExpirationTime(*accesstoken.AccessToken) *errors.RestError
}

type repository struct {
	session *gocql.Session
}

// NewDBRepository return new db repository
func NewDBRepository(session *gocql.Session) Repository {
	return &repository{
		session: session,
	}
}

func (r *repository) GetByID(accessTokenID string) (*accesstoken.AccessToken, *errors.RestError) {
	result := &accesstoken.AccessToken{}
	if err := r.session.Query(queryGetAccessToken, accessTokenID).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expired,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewInternalServerError("No access token found with a given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return result, nil
}

func (r *repository) Create(token *accesstoken.AccessToken) *errors.RestError {
	if err := r.session.Query(
		queryCreateAccessToken,
		token.AccessToken,
		token.UserID,
		token.ClientID,
		token.Expired,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *repository) UpdateExpirationTime(token *accesstoken.AccessToken) *errors.RestError {
	if err := r.session.Query(
		queryUpdateExpirationTime,
		token.Expired,
		token.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
