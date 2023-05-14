package handlers

import (
	"errors"
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
	// check if user already exists by login or username
	var count int
	if err := inits.Session.Query("SELECT COUNT(*) FROM users WHERE login=? ALLOW FILTERING",
		user.Login).Scan(&count); err != nil {
		return err
	}
	if err := inits.Session.Query("SELECT COUNT(*) FROM users WHERE email=? ALLOW FILTERING",
		user.Email).Scan(&count); err != nil {
		return err
	}

	if count > 0 {
		return errors.New("user already exists")
	}

	if err := inits.Session.Query("INSERT INTO users (user_id, email, login, password) VALUES (?, ?, ?, ?)",
		gocql.TimeUUID(), user.Email, user.Login, user.Password).Exec(); err != nil {
		return err
	}
	return nil
}

func LoginUser(user structs.LoginUser) (structs.User, error) {
	var dbUser structs.User
	err := inits.Session.Query("SELECT * FROM users WHERE email = ? AND password = ? ALLOW FILTERING",
		user.Login, user.Password).Scan(&dbUser.User_id, &dbUser.Email, &dbUser.Login, &dbUser.Password)
	if err != nil {
		return structs.User{}, errors.New("invalid credentials")
	}
	return dbUser, nil
}
