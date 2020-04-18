package accesstoken

import (
	"testing"
	"time"
)

func TestGetNewAccessToken_ReturnInitValue(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("Brand new access token should not be expired.")
	}

	if at.AccessToken != "" {
		t.Error("New Access token should not have access token")
	}

	if at.ClientID != 0 {
		t.Error("New Access token should not have user id")
	}

	if at.UserID != 0 {
		t.Error("New Access token should not have user id")
	}
}

func TestIsExpired_TokenExpiredWhenEmpty(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired() {
		t.Error("Empty access token should be expired.")
	}
}

func TestIsExpired_TokenIsNotExpired(t *testing.T) {
	at := AccessToken{}
	at.Expired = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("Access token expired 3 hours from now should not be expired")
	}
}

func TestIsExpired_TokenIsExpired(t *testing.T) {
	at := AccessToken{}
	at.Expired = time.Now().UTC().Add(-time.Hour).Unix()
	if !at.IsExpired() {
		t.Error("Access token should expired 1 hours")
	}
}
