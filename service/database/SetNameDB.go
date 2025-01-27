package database

func (db *appdbimpl) SetName(id string, newname string) error {

	_, err := db.c.Exec("UPDATE users SET userName = ? WHERE userId = ?", newname, id)
	if err != nil {
		return err
	}

	return err
}
