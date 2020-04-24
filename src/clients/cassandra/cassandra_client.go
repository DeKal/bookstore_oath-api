package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	godotenv.Load()

	cluster = gocql.NewCluster(getHost())
	cluster.Keyspace = getKeySpace()
	cluster.Consistency = gocql.Quorum
}

// NewCassandraSession create new session from cluster
func NewCassandraSession() *gocql.Session {
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return session
}
