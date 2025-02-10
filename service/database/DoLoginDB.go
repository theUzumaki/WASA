package database

import (
	"database/sql"
	"errors"
	"strconv"
)

func (db *appdbimpl) LoginManager(user User) (User, string, error) {

	row := db.c.QueryRow("SELECT * FROM users WHERE userName = ?", user.Name)

	var newuser User
	var id int
	err := row.Scan(&newuser.Id, &newuser.Name)
	if !errors.Is(err, sql.ErrNoRows) {
		return newuser, "user exist", nil
	} else if errors.Is(err, sql.ErrNoRows) {
		newuser.Name = user.Name
	} else if err != nil {
		return newuser, "", err
	}

	row = db.c.QueryRow("SELECT MAX(userId) FROM users")
	var stringId string
	err = row.Scan(&stringId)
	if stringId == "" {
		id = 0
	} else if err != nil {
		return newuser, "", err
	} else {
		id, err = strconv.Atoi(stringId)
		if err != nil {
			return newuser, "", err
		}
	}
	newuser.Id = id + 1

	_, err = db.c.Exec("INSERT INTO users (userId, userName) VALUES (?, ?);", newuser.Id, newuser.Name)
	if err != nil {
		return newuser, "", err
	}

	return newuser, "", nil
}
