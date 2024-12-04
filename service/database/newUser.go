package database

func (db *appdbimpl) newUser(user User) error {

	row := db.c.QueryRow("SELECT MAX(userId) FROM user")

	var id int
	row.Scan(&id, nil, nil, nil)

	user.Id = id + 1
	db.c.Exec("INSERT INTO user (userId, userName) VALUES (?, ?)", user.Id, user.Name)

	return nil
}
