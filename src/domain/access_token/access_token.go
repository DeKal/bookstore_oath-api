package accesstoken

import "time"

const (
	expirationTime = 24
)

// AccessToken to verify user
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expired     int64  `json:"expired"`
}

// GetNewAccessToken Get new Access Token with Expired
func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expired: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired check if access token is expired or not
func (at AccessToken) IsExpired() bool {
	if at.Expired == 0 {
		return true
	}
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expired, 0)

	return expirationTime.Before(now)
}
