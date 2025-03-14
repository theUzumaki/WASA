package database

import "wasatext/service/structs"

func (db *appdbimpl) GetUserId(userName string) (structs.User, error) {

	row := db.c.QueryRow(`SELECT userId  
			FROM users
			WHERE userName = ?`, userName)

	var user structs.User
	err := row.Scan(&user.Id)
	if err != nil {
		return user, err
	}

	return user, err
}
