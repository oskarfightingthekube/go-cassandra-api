package inits

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func InitDB() {
	var err error
	cluster := gocql.NewCluster("185.180.205.124")
	cluster.Keyspace = "schools"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}
