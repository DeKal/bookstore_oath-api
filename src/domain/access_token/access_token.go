package accesstoken

import (
	"fmt"
	"strings"
	"time"

	"github.com/DeKal/bookstore_utils-go/crypto"
	"github.com/DeKal/bookstore_utils-go/errors"
)

const (
	expirationTime            = 24
	grantTypePassword         = "password"
	grantTypeCliendCredential = "client_credentials"
)

// AccessToken to verify user
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id,omitempty"`
	Expired     int64  `json:"expired"`
}

// Request for new user
type Request struct {
	GrantType    string `json:"grant_type"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

// Validate data from access token
func (req *Request) Validate() *errors.RestError {
	switch req.GrantType {
	case grantTypePassword:
		break
	case grantTypeCliendCredential:
		break
	default:
		return errors.NewBadRequestError("Wrong grant_type request.")
	}

	return nil
}

// GetNewAccessToken Get new Access Token with Expired
func GetNewAccessToken(userID int64) *AccessToken {
	return &AccessToken{
		UserID:  userID,
		Expired: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired check if access token is expired or not
func (at *AccessToken) IsExpired() bool {
	if at.Expired == 0 {
		return true
	}
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expired, 0)

	return expirationTime.Before(now)
}

// Validate data from access token
func (at *AccessToken) Validate() *errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("Invalid access token id")
	}

	if at.UserID <= 0 {
		return errors.NewBadRequestError("Invalid user id")
	}

	if at.ClientID <= 0 {
		return errors.NewBadRequestError("Invalid client id")
	}

	if at.Expired <= 0 {
		return errors.NewBadRequestError("Invalid expired time")
	}

	return nil
}

// Generate generate new access token
func (at *AccessToken) Generate() {
	at.AccessToken = crypto.GetMD5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expired))
}
