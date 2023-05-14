package main

import (
	"go-cassandra-api/inits"
	"go-cassandra-api/routes"
)

func main() {
	routes.Run()
	inits.InitDB()
}
