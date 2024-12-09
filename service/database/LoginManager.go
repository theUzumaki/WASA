package database

func (db *appdbimpl) LoginManager(user User) (string, error) {

	row := db.c.QueryRow("SELECT * FROM users WHERE userName = ?", user.Name)
	if row != nil {
		return "user exist", nil
	}

	row = db.c.QueryRow("SELECT MAX(userId) FROM users")
	var id int
	row.Scan(&id)
	user.Id = id + 1
	db.c.Exec("INSERT INTO users (userId, userName) VALUES (?, ?)", user.Id, user.Name)

	return "", nil
}
