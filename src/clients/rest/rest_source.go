package rest

import "os"

const (
	host = "users_host"
)

func getUsersRestHost() string {
	return os.Getenv(host)
}
