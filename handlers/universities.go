package handlers

import (
	"go-cassandra-api/inits"
	"go-cassandra-api/structs"

	"github.com/gocql/gocql"
)

func GetUniversities() ([]structs.University, error) {
	var universities []structs.University
	m := map[string]interface{}{}
	iter := inits.Session.Query("SELECT * FROM university").Iter()
	for iter.MapScan(m) {
		universities = append(universities, structs.University{
			University_id: m["university_id"].(gocql.UUID).String(),
			City:          m["city"].(string),
			Country:       m["country"].(string),
			Name:          m["name"].(string),
		})
		m = map[string]interface{}{}
	}

	return universities, nil
}
