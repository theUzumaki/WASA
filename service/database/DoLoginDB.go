package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) LoginManager(user User) (int, string, error) {

	row := db.c.QueryRow("SELECT * FROM users WHERE userName = ?", user.Name)

	var id int
	err := row.Scan(&id, nil)
	if !errors.Is(err, sql.ErrNoRows) {
		return id, "user exist", nil
	}

	row = db.c.QueryRow("SELECT MAX(userId) FROM users")
	var stringId string
	err = row.Scan(&stringId)
	if stringId == "" {
		id = 1
	} else if err != nil {
		return -1, "", err
	}
	user.Id = id + 1

	_, err = db.c.Exec("INSERT INTO users (userId, userName) VALUES (?, ?);", user.Id, user.Name)
	if err != nil {
		return -1, "", err
	}

	return user.Id, "", nil
}
