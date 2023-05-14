package handlers

import (
	"go-cassandra-api/inits"
	"go-cassandra-api/structs"

	"github.com/gocql/gocql"
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

func GetUser(id string) (structs.User, error) {
	var user structs.User
	m := map[string]interface{}{}
	iter := inits.Session.Query("SELECT * FROM users WHERE user_id=? ALLOW FILTERING", id).Iter()
	for iter.MapScan(m) {
		user = structs.User{
			User_id:  m["user_id"].(gocql.UUID).String(),
			Email:    m["email"].(string),
			Login:    m["login"].(string),
			Password: HashPassword(m["password"].(string)),
		}
	}
	return user, nil
}

func AddUser(user structs.AddUser) error {
	if err := inits.Session.Query("INSERT INTO users (user_id, email, login, password) VALUES (?, ?, ?, ?)",
		gocql.TimeUUID(), user.Email, user.Login, user.Password).Exec(); err != nil {
		return err
	}
	return nil
}
