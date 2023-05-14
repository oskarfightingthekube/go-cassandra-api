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

func GetMajors() ([]structs.MajorsWithUniversity, error) {
	var majors []structs.MajorsWithUniversity
	m := map[string]interface{}{}
	iter := inits.Session.Query("SELECT * FROM majors").Iter()
	for iter.MapScan(m) {
		majors = append(majors, structs.MajorsWithUniversity{
			Major_id:        m["major_id"].(gocql.UUID).String(),
			Name:            m["name"].(string),
			Category:        m["type"].(string),
			University_id:   m["university_id"].(gocql.UUID).String(),
			University_name: "",
		})
		m = map[string]interface{}{}
	}

	// Get university names based on university IDs
	for i, major := range majors {
		var university_name string
		err := inits.Session.Query("SELECT name FROM university WHERE university_id = ?", major.University_id).Scan(&university_name)
		if err != nil {
			return nil, err
		}
		majors[i].University_name = university_name
	}

	return majors, nil
}
