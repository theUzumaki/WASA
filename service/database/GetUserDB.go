package database

func (db *appdbimpl) GetUserId(userName string) (User, error) {

	row := db.c.QueryRow(`SELECT userId  
			FROM users
			WHERE userName = ?`, userName)

	var user User
	err := row.Scan(&user.Id)
	if err != nil {
		return user, err
	}

	return user, err
}
