package database

func (db *appdbimpl) SetMyPhoto(id string, newpicture string) error {

	_, err := db.c.Exec("UPDATE users SET picture = ? WHERE userId = ?", newpicture, id)
	if err != nil {
		return err
	}

	return err
}
