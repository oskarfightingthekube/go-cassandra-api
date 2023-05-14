package handlers

import (
	"errors"
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

func GetMajorByName(name string) ([]structs.MajorsWithUniversity, error) {
	var majors []structs.MajorsWithUniversity
	m := map[string]interface{}{}
	iter := inits.Session.Query("SELECT * FROM majors WHERE name = ? ALLOW FILTERING", name).Iter()
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

func GetMajorByType(category string) ([]structs.MajorsWithUniversity, error) {
	var majors []structs.MajorsWithUniversity
	m := map[string]interface{}{}
	iter := inits.Session.Query("SELECT * FROM majors WHERE type = ? ALLOW FILTERING", category).Iter()
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

func GetDepartments() ([]structs.DepartmentWithUniversity, error) {
	var departments []structs.DepartmentWithUniversity
	m := map[string]interface{}{}
	iter := inits.Session.Query("SELECT * FROM departments").Iter()
	for iter.MapScan(m) {
		departments = append(departments, structs.DepartmentWithUniversity{
			Department_id:   m["department_id"].(gocql.UUID).String(),
			Name:            m["name"].(string),
			University_id:   m["university_id"].(gocql.UUID).String(),
			University_name: "",
		})
		m = map[string]interface{}{}
	}

	// Get university names based on university IDs
	for i, department := range departments {
		var university_name string
		err := inits.Session.Query("SELECT name FROM university WHERE university_id = ?", department.University_id).Scan(&university_name)
		if err != nil {
			return nil, err
		}
		departments[i].University_name = university_name
	}

	return departments, nil
}

func GetDepartmentByUniversity(name string) ([]structs.DepartmentWithUniversity, error) {
	var universityID gocql.UUID
	err := inits.Session.Query("SELECT university_id FROM university WHERE name = ? ALLOW FILTERING", name).Scan(&universityID)
	if err != nil {
		return nil, err
	}
	var departments []structs.DepartmentWithUniversity
	m := map[string]interface{}{}
	iter := inits.Session.Query("SELECT * FROM departments WHERE university_id = ? ALLOW FILTERING", universityID).Iter()
	for iter.MapScan(m) {
		departments = append(departments, structs.DepartmentWithUniversity{
			Department_id:   m["department_id"].(gocql.UUID).String(),
			Name:            m["name"].(string),
			University_id:   m["university_id"].(gocql.UUID).String(),
			University_name: name,
		})
		m = map[string]interface{}{}
	}

	return departments, nil
}

func AddUniversity(university structs.AddUniversity) error {
	var count int
	if err := inits.Session.Query("SELECT COUNT(*) FROM university WHERE name=? ALLOW FILTERING",
		university.Name).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("university already exists")
	}
	if err := inits.Session.Query("INSERT INTO university (university_id, name, country, city) VALUES (?, ?, ?, ?)",
		gocql.TimeUUID(), university.Name, university.Country, university.City).Exec(); err != nil {
		return err
	}
	return nil
}
