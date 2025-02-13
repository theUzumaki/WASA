package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) SetName(id string, newname string) error {

	row := db.c.QueryRow("SELECT * FROM users WHERE userName = ?", newname)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Picture)
	if errors.Is(err, sql.ErrNoRows) {
		_, err = db.c.Exec("UPDATE users SET userName = ? WHERE userId = ?", newname, id)
		if err != nil {
			return err
		}
	} else {
		log.Println("CHECK")
		return errors.New("already taken")
	}

	return err
}
