package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

var (
	cluster *gocql.ClusterConfig
	session *gocql.Session
)

func init() {
	godotenv.Load()

	cluster = gocql.NewCluster(getHost())
	cluster.Keyspace = getKeySpace()
	cluster.Consistency = gocql.Quorum

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}

// GetSession return Session of Cassandra
func GetSession() *gocql.Session {
	return session
}
