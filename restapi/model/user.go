package model

import (
	"errors"
	"restapi/db"
	"restapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:required`
	Password string `binding:required`
}

var users = []User{}

func (u User) Save() error {
	query := `insert into users(email,password) values (?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hash)
	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()

	u.ID = userId
	users = append(users, u)

	return err
}

func (u *User) ValidateCredentials() error {
	query := `select id, password from users where email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrivePassword string
	err := row.Scan(&u.ID, &retrivePassword)

	if err != nil {
		return err
	}

	passwordIsvalid := utils.CheckPasswordHash(u.Password, retrivePassword)

	if !passwordIsvalid {
		return errors.New("Password Invalid")
	}
	return nil
}
