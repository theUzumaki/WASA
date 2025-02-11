package database

func (db *appdbimpl) GetUsers(userName string) ([]User, error) {

	rowsUsers, err := db.c.Query(`SELECT *  
			FROM users
			WHERE userName LIKE ?`, userName+"%")
	if err != nil {
		return nil, err
	}

	var users []User
	defer rowsUsers.Close()
	for rowsUsers.Next() {
		var user User
		err = rowsUsers.Scan(&user.Id, &user.Name, &user.Picture)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}
