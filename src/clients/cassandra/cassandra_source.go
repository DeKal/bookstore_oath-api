package cassandra

import "os"

const (
	host     = "cassandra_host"
	keyspace = "cassandra_keyspace"
)

func getHost() string {
	return os.Getenv(host)
}

func getKeySpace() string {
	return os.Getenv(keyspace)
}
