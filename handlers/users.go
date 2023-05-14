package handlers

import (
	"go-cassandra-api/inits"
	"go-cassandra-api/structs"

	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() ([]structs.User, error) {
	var users []structs.User
	m := map[string]interface{}{}
	iter := inits.Session.Query("SELECT * FROM users").Iter()
	for iter.MapScan(m) {
		users = append(users, structs.User{
			User_id:  m["user_id"].(gocql.UUID).String(),
			Email:    m["email"].(string),
			Login:    m["login"].(string),
			Password: HashPassword(m["password"].(string)),
		})
		m = map[string]interface{}{}
	}
	return users, nil
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 1)
	return string(bytes)
}
