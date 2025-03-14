package database

import "wasatext/service/structs"

func (db *appdbimpl) GetUsers(userName string) ([]structs.User, error) {

	rowsUsers, err := db.c.Query(`SELECT *  
			FROM users
			WHERE userName LIKE ?`, userName+"%")
	if err != nil {
		return nil, err
	}

	var users []structs.User
	for rowsUsers.Next() {
		var user structs.User
		if rowsUsers.Err() != nil {
			return nil, err
		}

		err = rowsUsers.Scan(&user.Id, &user.Name, &user.Picture)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rowsUsers.Close()
	if err != nil {
		return users, err
	}

	return users, err
}
