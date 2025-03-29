package database

import (
	"database/sql"
	"errors"
	"wasatext/service/structs"
)

func (db *appdbimpl) SetName(id string, newname string) error {

	row := db.c.QueryRow("SELECT * FROM users WHERE userName = ? AND userId <> ?", newname, id)

	var user structs.User
	err := row.Scan(&user.Id, &user.Name, &user.Picture)
	if errors.Is(err, sql.ErrNoRows) {
		_, err = db.c.Exec("UPDATE users SET userName = ? WHERE userId = ?", newname, id)
		if err != nil {
			return err
		}
	} else {
		return errors.New("already taken")
	}

	return err
}
