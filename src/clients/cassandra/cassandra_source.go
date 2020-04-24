package cassandra

import "os"

const (
	host     = "cassandra_host"
	keyspace = "cassandra_keyspace"
)

func getHost() string {
	return os.Getenv("cassandra_host")
}

func getKeySpace() string {
	return os.Getenv("cassandra_keyspace")
}
